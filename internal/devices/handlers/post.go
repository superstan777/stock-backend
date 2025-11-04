package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices"
	"github.com/superstan777/stock-backend/internal/devices/repository"
)

func CreateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var d devices.Device
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := repository.CreateDevice(db.DB, &d); err != nil {
		http.Error(w, "DB insert error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}