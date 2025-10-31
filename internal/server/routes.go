package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/handlers"
)

func (s *Server) routes() {

	s.Router.Use(LoggingMiddleware)
	s.Router.Use(CORSMiddleware)

	s.Router.Route("/api", func(r chi.Router) {
		r.Get("/health", handlers.HealthHandler)
		r.Get("/users", handlers.UsersHandler)
	})
}