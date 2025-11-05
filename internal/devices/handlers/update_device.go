package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

// UpdateDeviceHandler obsługuje PUT /api/device/{id}
func UpdateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing device ID")
		return
	}

	var d devices.Device
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		apiresponse.JSONError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	if err := repository.UpdateDevice(db.DB, id, &d); err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database update error: "+err.Error())
		return
	}

	apiresponse.JSONSuccess(w, http.StatusOK, "Device updated successfully", map[string]string{
		"id":     id,
		"status": "updated",
	})
}