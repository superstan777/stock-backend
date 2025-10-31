package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/handlers"
)

func (s *Server) routes() {
	s.router.Use(LoggingMiddleware)
	s.router.Use(CORSMiddleware)

	s.router.Route("/api", func(r chi.Router) {
		r.Get("/health", handlers.HealthHandler)
		r.Get("/users", handlers.UsersHandler)
	})
}