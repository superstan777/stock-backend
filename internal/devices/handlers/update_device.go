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

// UpdateDeviceHandler godoc
// @Summary      Update a device
// @Description  Updates an existing device in the database by ID
// @Tags         devices
// @Accept       json
// @Produce      json
// @Param        id     path      string             true  "Device ID"
// @Param        device body      devices.Device     true  "Updated device data"
// @Success      200  {object}  apiresponse.UpdatedResponse
// @Failure      400  {object}  apiresponse.ErrorResponse
// @Failure      500  {object}  apiresponse.ErrorResponse
// @Router       /devices/{id} [put]
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

	apiresponse.JSONUpdated(w, d)
}