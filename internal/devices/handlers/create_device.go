package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func CreateDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var d devices.Device

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		apiresponse.JSONError(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	if d.DeviceType == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Device type is required")
		return
	}

	if err := repository.CreateDevice(db.DB, &d); err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database insert error: "+err.Error())
		return
	}

	apiresponse.JSONCreated(w, d)
}