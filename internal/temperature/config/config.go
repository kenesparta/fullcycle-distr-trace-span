package config

type Config struct {
	Temperature Temperature
	CEP         CEP
	App         App
}

type Temperature struct {
	ApiKey string
	URL    string
}

type CEP struct {
	URL string
}

type App struct {
	Port string
}
