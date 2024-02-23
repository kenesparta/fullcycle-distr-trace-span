package service

import (
	"context"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/dto"
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

	wRes, err := ts.Weather.Execute(ctx, location)
	if err != nil {
		return nil, err
	}

	return &pb.TemperatureResponse{
		Location: wRes.Location,
		TempC:    float32(wRes.TempC),
		TempF:    float32(wRes.TempF),
		TempK:    float32(wRes.TempK),
	}, nil
}
