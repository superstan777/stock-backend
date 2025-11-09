package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/relations/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func HasActiveRelationHandler(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "device_id")
	if deviceID == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing device_id")
		return
	}

	active, err := repository.HasActiveRelation(db.DB, deviceID)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	apiresponse.JSONHasActiveRelation(w, http.StatusOK, active)
}