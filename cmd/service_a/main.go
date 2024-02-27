package main

import (
	"log"
	"net/http"

	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/grpc"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/temperature/config"
)

func main() {
	var cfg config.Config
	viperCfg := config.NewViper("env.json")
	viperCfg.ReadViper(&cfg)

	genRt := grpc.SetClient(cfg)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /temperature", genRt.Temperature)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println("server error", err)
		return
	}
}
