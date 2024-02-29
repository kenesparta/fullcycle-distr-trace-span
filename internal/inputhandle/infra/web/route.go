package web

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/dto"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func (gr *Server) temperature(writer http.ResponseWriter, request *http.Request) {
	carrier := propagation.HeaderCarrier(request.Header)
	ctx := request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanFn := gr.TemplateData.OTELTracer.Start(ctx, gr.TemplateData.RequestNameOtel)
	defer spanFn.End()
	time.Sleep(time.Millisecond * 200)

	bodyBytes, readErr := io.ReadAll(request.Body)
	if readErr != nil {
		http.Error(writer, readErr.Error(), http.StatusBadRequest)
		return
	}

	var location dto.LocationInput
	if unmErr := json.Unmarshal(bodyBytes, &location); unmErr != nil {
		http.Error(writer, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	if err := entity.CEPValidation(location.CEP); err != nil {
		http.Error(writer, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	reqCtx, reqCtxErr := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://service_b:50055/temperature?cep="+location.CEP,
		nil,
	)
	if reqCtxErr != nil {
		http.Error(writer, reqCtxErr.Error(), http.StatusInternalServerError)
		return
	}

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(reqCtx.Header))
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	clientDo, doErr := httpClient.Do(reqCtx)
	if doErr != nil {
		http.Error(writer, doErr.Error(), http.StatusInternalServerError)
		return
	}
	defer clientDo.Body.Close()

	switch clientDo.StatusCode {
	case http.StatusUnprocessableEntity:
		http.Error(writer, entity.ErrCEPNotValid.Error(), http.StatusUnprocessableEntity)
		return
	case http.StatusNotFound:
		http.Error(writer, entity.ErrCEPNotFound.Error(), http.StatusNotFound)
		return
	}

	locRespBody, bReadErr := io.ReadAll(clientDo.Body)
	if bReadErr != nil {
		http.Error(writer, bReadErr.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	var locTempResp dto.TemperatureAPIOutput
	if err := json.Unmarshal(locRespBody, &locTempResp); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

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
