package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

// DeleteDeviceHandler obsługuje DELETE /api/device/{id}
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

	apiresponse.JSONSuccess(w, http.StatusOK, "Device deleted successfully", map[string]string{
		"id":     id,
		"status": "deleted",
	})
}