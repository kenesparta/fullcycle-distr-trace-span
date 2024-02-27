package grpc

//go:generate protoc --go_out=. --go-grpc_out=. ./protofiles/temperature.proto
//go:generate protoc --go_out=. --go-grpc_out=. ./protofiles/health.proto

import (
	_ "github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/grpc/pb"
	_ "github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/infra/grpc/service"
)
