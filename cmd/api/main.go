package main

import (
	"log"
	"net/http"

	"github.com/superstan777/stock-backend/internal/server"
)

func main() {
	srv := server.NewServer()

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", srv.Router); err != nil {
		log.Fatal(err)
	}
}