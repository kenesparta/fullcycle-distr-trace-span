package config

type Config struct {
	ServiceA    ServiceA
	ServiceB    ServiceB
	Temperature Temperature
	CEP         CEP
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

type ServiceB struct {
	Port string
	Host string
}

type ServiceA struct {
	Port string
}

type Zipkin struct {
	Host     string
	Endpoint string
}

type Otel struct {
	Host string
}
