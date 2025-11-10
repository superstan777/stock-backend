package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

// GetDeviceHandler godoc
// @Summary Get a device by ID
// @Description Retrieve detailed information about a specific device using its ID
// @Tags devices
// @Accept json
// @Produce json
// @Param id path string true "Device ID"
// @Success 200 {object} apiresponse.Response "Response with device data in 'data' field (devices.Device)"
// @Failure 400 {object} apiresponse.ErrorResponse
// @Failure 404 {object} apiresponse.ErrorResponse
// @Failure 500 {object} apiresponse.ErrorResponse
// @Router /devices/{id} [get]
func GetDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing device ID")
		return
	}

	device, err := repository.GetDeviceByID(db.DB, id)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	if device == nil {
		apiresponse.JSONError(w, http.StatusNotFound, "Device not found")
		return
	}

	apiresponse.JSONSuccess(w, http.StatusOK, device, nil)
}