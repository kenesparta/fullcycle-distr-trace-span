package service

import (
	"context"
	"errors"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/dto"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/entity"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/grpc/pb"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/usecase"
)

type TemperatureService struct {
	pb.TemperatureServiceServer
	Weather usecase.GetWeather
}

func NewTemperatureService(getWeather usecase.GetWeather) TemperatureService {
	return TemperatureService{
		Weather: getWeather,
	}
}

func (ts TemperatureService) GetWeather(
	ctx context.Context,
	in *pb.Location,
) (*pb.TemperatureResponse, error) {
	location := dto.LocationInput{
		CEP: in.Cep,
	}

	wRes, execErr := ts.Weather.Execute(ctx, location)
	switch {
	case errors.Is(execErr, entity.ErrCEPNotFound):
		return &pb.TemperatureResponse{
			TempError: pb.ErrorCode_CEP_NOT_FOUND,
		}, nil
	case errors.Is(execErr, entity.ErrCEPNotValid):
		return &pb.TemperatureResponse{
			TempError: pb.ErrorCode_INVALID_CEP,
		}, nil
	}
	if execErr != nil {
		return nil, execErr
	}

	return &pb.TemperatureResponse{
		Location: wRes.Location,
		TempC:    float32(wRes.TempC),
		TempF:    float32(wRes.TempF),
		TempK:    float32(wRes.TempK),
	}, nil
}
