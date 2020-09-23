package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mikestefanello/surveys-microservices/vote-service/handler"
	"github.com/mikestefanello/surveys-microservices/vote-service/middleware"
)

// NewRouter creates a new router
func NewRouter(h *handler.VoteHTTPHandler) *chi.Mux {
	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// Vote routes
	r.Route("/vote", func(r chi.Router) {
		r.Use(middleware.AddSerializer)
		r.Post("/", h.Vote)
	})

	return r
}
