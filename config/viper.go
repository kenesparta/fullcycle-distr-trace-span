package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const fileExtension = "json"

type Viper struct {
	fileName string
}

func NewViper(fileName string) *Viper {
	return &Viper{
		fileName: fileName,
	}
}

func (v *Viper) ReadViper(config *Config) {
	viper.SetConfigName(v.fileName)
	viper.SetConfigType(fileExtension)
	viper.AddConfigPath(".")
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	v.readConfig(config)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		v.readConfig(config)
	})
}

func (v *Viper) readConfig(c *Config) {
	c.CEP.URL = viper.GetString("cep.url")
	c.Temperature.URL = viper.GetString("temperature.url")
	c.Temperature.ApiKey = viper.GetString("temperature.api_key")

	c.ServiceA.Port = viper.GetString("service_a.port")

	c.ServiceB.Port = viper.GetString("service_b.port")
	c.ServiceB.Host = viper.GetString("service_b.host")

	c.Zipkin.Host = viper.GetString("zipkin.host")
	c.Zipkin.Endpoint = viper.GetString("zipkin.endpoint")
}
