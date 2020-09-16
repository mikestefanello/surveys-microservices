package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mikestefanello/surveys-microservices/survey-service/handler"
	"github.com/mikestefanello/surveys-microservices/survey-service/middleware"
)

// NewRouter creates a new router
func NewRouter(h *handler.SurveyHTTPHandler) *chi.Mux {
	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// Survey routes
	r.Route("/surveys", func(r chi.Router) {
		r.Use(middleware.AddSerializer)
		r.Get("/", h.Collection)
		r.Get("/{id}", h.Get)
		r.Post("/", h.Post)
	})

	return r
}
