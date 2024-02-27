package service

import (
	"context"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/grpc/pb_health"
)

type HealthServer struct {
	pb_health.UnimplementedHealthServer
}

func (s *HealthServer) Check(
	_ context.Context,
	_ *pb_health.HealthCheckRequest,
) (*pb_health.HealthCheckResponse, error) {
	return &pb_health.HealthCheckResponse{
		Status: pb_health.HealthCheckResponse_SERVING,
	}, nil
}
