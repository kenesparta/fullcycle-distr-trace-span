package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/kenesparta/fullcycle-distr-trace-span/config"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/opentel"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/dto"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/entity"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/api"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	var cfg config.Config
	viperCfg := config.NewViper("env.json")
	viperCfg.ReadViper(&cfg)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	providerShutdown, provErr := opentel.InitProvider(
		"service_b",
		cfg.Zipkin.Endpoint,
	)
	if provErr != nil {
		return
	}

	gw := usecase.NewGetWeather(
		api.NewCEPFromAPI(&cfg),
		api.NewWeatherFromAPI(&cfg),
	)

	defer func() {
		if err := providerShutdown(ctx); err != nil {
			log.Printf("failed shuting down the tracer provider %s\n", err.Error())
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc(
		"GET /temperature",
		func(w http.ResponseWriter, r *http.Request) {
			carrier := propagation.HeaderCarrier(r.Header)
			hCtx := r.Context()
			hCtx = otel.GetTextMapPropagator().Extract(hCtx, carrier)

			tracer := otel.Tracer("service_b")
			_, span := tracer.Start(hCtx, "service_b:all")
			defer span.End()

			temperature, execErr := gw.Execute(hCtx, dto.LocationInput{
				CEP: r.URL.Query().Get("cep"),
			})

			switch {
			case errors.Is(execErr, entity.ErrCEPNotValid):
				http.Error(w, entity.ErrCEPNotValid.Error(), http.StatusUnprocessableEntity)
				return
			case errors.Is(execErr, entity.ErrCEPNotFound):
				http.Error(w, entity.ErrCEPNotFound.Error(), http.StatusNotFound)
				return
			case execErr != nil:
				http.Error(w, execErr.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(temperature); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
		},
	)

	if err := http.ListenAndServe(":"+cfg.ServiceB.Port, mux); err != nil {
		return
	}
}
