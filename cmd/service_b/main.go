package main

import (
	"fmt"
	"log"
	"net"

	conf "github.com/kenesparta/fullcycle-distr-trace-span/config"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/api"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/grpc/pb"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/grpc/pb_health"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/grpc/service"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var cfg conf.Config
	viperCfg := conf.NewViper("env.json")
	viperCfg.ReadViper(&cfg)

	getWeather := usecase.NewGetWeather(
		api.NewCEPFromAPI(&cfg),
		api.NewWeatherFromAPI(&cfg),
	)

	grpcServer := grpc.NewServer()
	pb.RegisterTemperatureServiceServer(
		grpcServer,
		service.NewTemperatureService(getWeather),
	)
	pb_health.RegisterHealthServer(grpcServer, &service.HealthServer{})
	reflection.Register(grpcServer)

	lis, listErr := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Grpc.Port))
	if listErr != nil {
		log.Println("error creating the TCP server")
		return
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error creating the TCP server")
		return
	}
}
