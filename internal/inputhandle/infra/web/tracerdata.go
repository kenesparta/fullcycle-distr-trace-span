package web

import (
	"go.opentelemetry.io/otel/trace"
	"time"
)

type TemplateData struct {
	Title              string
	ResponseTime       time.Duration
	BackgroundColor    string
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOtel    string
	OTELTracer         trace.Tracer
}
