package main

import (
	"log"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/server"
)

func main() {
    if err := db.Connect(); err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    srv := server.NewServer() // ten router ma już /api/users i /api/health z middleware

    log.Println("Server running on :8080")
    if err := http.ListenAndServe(":8080", srv.Router); err != nil {
        log.Fatal(err)
    }
}