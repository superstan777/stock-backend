package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

// CreateDeviceHandler godoc
// @Summary Create a new device
// @Description Create a new device record in the database
// @Tags devices
// @Accept json
// @Produce json
// @Param device body devices.Device true "Device object"
// @Success 201 {object} apiresponse.CreatedResponse
// @Failure 400 {object} apiresponse.ErrorResponse
// @Failure 500 {object} apiresponse.ErrorResponse
// @Router /devices [post]
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