package route

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/dto"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/grpc/pb"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/entity"
)

type GeneralRoute struct {
	GrpcServer pb.TemperatureServiceClient
}

func (gr *GeneralRoute) Temperature(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Error reading request body", http.StatusBadRequest)
		return
	}

	var location dto.LocationInput
	if unmErr := json.Unmarshal(bodyBytes, &location); unmErr != nil {
		http.Error(writer, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(request.Context(), 3*time.Second)
	defer cancel()
	r, getErr := gr.GrpcServer.GetWeather(ctx, &pb.Location{Cep: location.CEP})
	if getErr != nil {
		log.Printf("could not get weather: %v\n", getErr)
		return
	}

	switch r.GetTempError() {
	case pb.ErrorCode_INVALID_CEP:
		http.Error(writer, entity.ErrCEPNotValid.Error(), http.StatusUnprocessableEntity)
		return
	case pb.ErrorCode_CEP_NOT_FOUND:
		http.Error(writer, entity.ErrCEPNotFound.Error(), http.StatusNotFound)
		return
	}

	response := dto.TemperatureOutput{
		City:  r.GetLocation(),
		TempC: float64(r.GetTempC()),
		TempF: float64(r.GetTempF()),
		TempK: float64(r.GetTempK()),
	}

	writer.Header().Set("Content-Type", "application/json")
	jsonData, marshErr := json.Marshal(response)
	if marshErr != nil {
		http.Error(writer, "Error generating JSON", http.StatusInternalServerError)
		return
	}

	if _, wErr := writer.Write(jsonData); wErr != nil {
		http.Error(writer, "Error writing", http.StatusInternalServerError)
		return
	}
}
