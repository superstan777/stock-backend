package main

import (
	"log"

	"github.com/superstan777/stock-backend/internal/server"
)

func main() {
	s := server.NewServer(8080)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}