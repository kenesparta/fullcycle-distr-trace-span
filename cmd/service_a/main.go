package main

import (
	"context"
	conf "github.com/kenesparta/fullcycle-distr-trace-span/config"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/opentel"
	"github.com/kenesparta/fullcycle-distr-trace-span/internal/inputhandle/infra/web"
	"go.opentelemetry.io/otel"
	"log"
	"os"
	"os/signal"
)

func main() {
	var cfg conf.Config
	viperCfg := conf.NewViper("env.json")
	viperCfg.ReadViper(&cfg)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	providerShutdown, provErr := opentel.InitProvider(
		"service_a_orchestration",
		cfg.Zipkin.Endpoint,
	)
	if provErr != nil {
		return
	}

	defer func() {
		if err := providerShutdown(ctx); err != nil {
			log.Printf("failed shuting down the tracer provider %s\n", err.Error())
		}
	}()

	server := web.Server{
		TemplateData: web.TemplateData{
			Title:           "Service A: Orchestration",
			ExternalCallURL: cfg.ServiceB.Host,
			RequestNameOtel: "service_a_span",
			OTELTracer:      otel.Tracer("service_a"),
		},
	}

	server.Execute()
}
