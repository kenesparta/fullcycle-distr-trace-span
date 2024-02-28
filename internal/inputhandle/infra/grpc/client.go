package grpc

import (
	"github.com/kenesparta/fullcycle-distr-trace-span/config"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/grpc/pb"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/route"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func SetClient(cfg config.Config) *route.GeneralRoute {
	conn, dialErr := grpc.Dial(
		cfg.GrpcClient.Host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if dialErr != nil {
		log.Printf("did not connect: %v\n", dialErr)
		return nil
	}

	return &route.GeneralRoute{
		GrpcServer: pb.NewTemperatureServiceClient(conn),
	}
}
