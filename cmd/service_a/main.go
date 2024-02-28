package main

import (
	"log"
	"net/http"

	conf "github.com/kenesparta/fullcycle-distr-trace-span/config"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/grpc"
)

func main() {
	var cfg conf.Config
	viperCfg := conf.NewViper("env.json")
	viperCfg.ReadViper(&cfg)

	genRt := grpc.SetClient(cfg)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /temperature", genRt.Temperature)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println("server error", err)
		return
	}
}
