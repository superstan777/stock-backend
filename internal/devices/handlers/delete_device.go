package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

// DeleteDeviceHandler godoc
// @Summary Delete a device
// @Description Delete a device by its ID
// @Tags devices
// @Param id path string true "Device ID"
// @Success 200 {object} apiresponse.DeletedResponse
// @Failure 400 {object} apiresponse.ErrorResponse
// @Failure 500 {object} apiresponse.ErrorResponse
// @Router /devices/{id} [delete]
func DeleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing device ID")
		return
	}

	if err := repository.DeleteDevice(db.DB, id); err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database delete error: "+err.Error())
		return
	}

	apiresponse.JSONDeleted(w, id)
}