package usecase

import (
	"context"

	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/dto"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/entity"
)

type GetWeather struct {
	locationRepo entity.LocationRepo
	tempRepo     entity.TemperatureRepo
}

func NewGetWeather(
	locationRepo entity.LocationRepo,
	tempRepo entity.TemperatureRepo,
) *GetWeather {
	return &GetWeather{
		locationRepo: locationRepo,
		tempRepo:     tempRepo,
	}
}

func (gw *GetWeather) Execute(
	ctx context.Context,
	input dto.LocationInput,
) (dto.TemperatureOutput, error) {
	if cepEr := entity.CEPValidation(input.CEP); cepEr != nil {
		return dto.TemperatureOutput{}, cepEr
	}

	location, err := gw.locationRepo.Get(ctx, input.CEP)
	if err != nil {
		return dto.TemperatureOutput{}, err
	}

	temperature, err := gw.tempRepo.Get(ctx, location.Localidade)
	if err != nil {
		return dto.TemperatureOutput{}, err
	}

	return dto.TemperatureOutput{
		Location: location.Localidade,
		TempC:    temperature.Celsius(),
		TempF:    temperature.Fahrenheit(),
		TempK:    temperature.Kelvin(),
	}, nil
}
