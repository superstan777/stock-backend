package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Server represents the main HTTP server.
type Server struct {
	router *chi.Mux
	port   int
}

// NewServer creates a new instance of Server.
func NewServer(port int) *Server {
	s := &Server{
		router: chi.NewRouter(),
		port:   port,
	}
	s.routes() // attach all routes and middleware
	return s
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	fmt.Printf("Server is running on port %d...\n", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}