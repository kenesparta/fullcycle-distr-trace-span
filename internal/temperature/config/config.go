package config

type Config struct {
	Temperature Temperature
	CEP         CEP
	Grpc        Grpc
	GrpcClient  GrpcClient
}

type Temperature struct {
	ApiKey string
	URL    string
}

type CEP struct {
	URL string
}

type Grpc struct {
	Port string
}

type GrpcClient struct {
	Host string
}
