package config

type Config struct {
	Temperature Temperature
	CEP         CEP
	Grpc        Grpc
	GrpcClient  GrpcClient
	API         API
	Zipkin      Zipkin
	Otel        Otel
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

type API struct {
	Port string
}

type Zipkin struct {
	Host string
}

type Otel struct {
	Host string
}
