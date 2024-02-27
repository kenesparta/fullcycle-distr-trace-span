package main

import (
	"context"
	"log"
	"time"

	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/grpc/pb"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var cfg config.Config
	viperCfg := config.NewViper("env.json")
	viperCfg.ReadViper(&cfg)

	conn, err := grpc.Dial(cfg.GrpcClient.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTemperatureServiceClient(conn)

	cep := "12345"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetWeather(ctx, &pb.Location{Cep: cep})
	if err != nil {
		log.Fatalf("could not get weather: %v", err)
	}
	log.Printf("Temperature in %s: %f°C, %f°F, %f°K", r.GetLocation(), r.GetTempC(), r.GetTempF(), r.GetTempK())
}
