package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/relations/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func GetRelationsByDeviceHandler(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "device_id")
	if deviceID == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing device_id")
		return
	}

	list, err := repository.GetByDevice(db.DB, deviceID)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	if len(list) == 0 {
		apiresponse.JSONSuccess(w, http.StatusOK, "No relations found for this device", []interface{}{})
		return
	}

	apiresponse.JSONSuccess(w, http.StatusOK, "Relations fetched successfully", list)
}