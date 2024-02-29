package web

import (
	"log"
	"net/http"
)

type Server struct {
	TemplateData TemplateData
	mux          *http.ServeMux
}

func (gr *Server) prepare() {
	gr.mux = http.NewServeMux()
	gr.mux.HandleFunc("POST /temperature", gr.temperature)
}

func (gr *Server) run() {
	if err := http.ListenAndServe(":8080", gr.mux); err != nil {
		log.Println("server error", err)
		return
	}
}

func (gr *Server) Execute() {
	gr.prepare()
	gr.run()
}
