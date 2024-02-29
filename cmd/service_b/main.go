package main

import (
	"encoding/json"
	"errors"
	"github.com/kenesparta/fullcycle-distr-trace-span/config"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/dto"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/entity"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/api"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/usecase"
	"net/http"
)

func main() {
	var cfg config.Config
	viperCfg := config.NewViper("env.json")
	viperCfg.ReadViper(&cfg)

	gw := usecase.NewGetWeather(
		api.NewCEPFromAPI(&cfg),
		api.NewWeatherFromAPI(&cfg),
	)

	mux := http.NewServeMux()
	mux.HandleFunc(
		"GET /temperature",
		func(w http.ResponseWriter, r *http.Request) {
			temperature, execErr := gw.Execute(r.Context(), dto.LocationInput{
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
