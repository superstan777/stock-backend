package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices"
	"github.com/superstan777/stock-backend/internal/devices/repository"
)

func UpdateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var d devices.Device
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := repository.UpdateDevice(db.DB, id, &d); err != nil {
		http.Error(w, "DB update error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}