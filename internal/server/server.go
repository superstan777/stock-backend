package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	s := &Server{
		Router: chi.NewRouter(),
	}

	// Ustawiamy trasy
	s.routes()

	return s
}

func (s *Server) Start(addr string) {
	log.Printf("Server running on %s", addr)
	if err := http.ListenAndServe(addr, s.Router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}