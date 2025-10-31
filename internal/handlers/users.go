package handlers

import (
	"encoding/json"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []map[string]string{
		{"id": "1", "name": "Alice"},
		{"id": "2", "name": "Bob"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}