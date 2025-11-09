package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

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