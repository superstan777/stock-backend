package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
)

func GetDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	device, err := repository.GetDeviceByID(db.DB, id)
	if err != nil {
		http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if device == nil {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}