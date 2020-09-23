package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mikestefanello/surveys-microservices/vote-service/config"
)

// NewHTTPServer creates a new HTTP server
func NewHTTPServer(router *chi.Mux, cfg config.HTTPConfig) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port),
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
