package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/dto"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func (gr *Server) temperature(writer http.ResponseWriter, request *http.Request) {
	carrier := propagation.HeaderCarrier{}
	ctx := request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanFn := gr.TemplateData.OTELTracer.Start(ctx, gr.TemplateData.RequestNameOtel)
	defer spanFn.End()

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

	response, err := http.Get("http://service_b:50055")
	if err != nil {
		http.Error(writer, "Error making the request", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusUnprocessableEntity:
		http.Error(writer, entity.ErrCEPNotValid.Error(), http.StatusUnprocessableEntity)
		return
	case http.StatusNotFound:
		http.Error(writer, entity.ErrCEPNotFound.Error(), http.StatusNotFound)
		return
	}

	locRespBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var locTempResp dto.TemperatureAPIOutput
	json.Unmarshal(locRespBody, &locTempResp)

	writer.Header().Set("Content-Type", "application/json")
	jsonData, marshErr := json.Marshal(dto.TemperatureOutput{
		City:  locTempResp.Location,
		TempC: locTempResp.TempC,
		TempF: locTempResp.TempF,
		TempK: locTempResp.TempK,
	})
	if marshErr != nil {
		http.Error(writer, "Error generating JSON", http.StatusInternalServerError)
		return
	}

	if _, wErr := writer.Write(jsonData); wErr != nil {
		http.Error(writer, "Error writing", http.StatusInternalServerError)
		return
	}
}
