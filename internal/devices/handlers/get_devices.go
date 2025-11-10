package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

// GetDevicesHandler godoc
// @Summary Get list of devices
// @Description Returns a paginated list of devices. Optionally filter by device type or query parameters.
// @Tags devices
// @Accept json
// @Produce json
// @Param device_type path string false "Device type (computers | monitors)"
// @Param page query int false "Page number for pagination (default: 1)"
// @Param filters query string false "Optional filters (e.g. model, order_id, install_status)"
// @Success 200 {object} apiresponse.PaginatedResponse
// @Failure 400 {object} apiresponse.ErrorResponse
// @Failure 500 {object} apiresponse.ErrorResponse
// @Router /devices [get]
// @Router /devices/{device_type} [get]
func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	deviceType := chi.URLParam(r, "device_type")
	query := r.URL.Query()

	// Dozwolone typy urządzeń
	validTypes := map[string]string{
		"computers": "computer",
		"monitors":  "monitor",
	}

	// Parsowanie numeru strony
	page := 1
	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		} else {
			apiresponse.JSONError(w, http.StatusBadRequest, "Invalid page number")
			return
		}
	}

	// Filtry (pomijamy `page`)
	filters := make(map[string]string)
	for key, values := range query {
		if key == "page" {
			continue
		}
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	// Walidacja typu urządzenia
	var singular string
	if deviceType != "" {
		var ok bool
		singular, ok = validTypes[deviceType]
		if !ok {
			apiresponse.JSONError(w, http.StatusBadRequest, "Invalid device type: "+deviceType)
			return
		}
	}

	// Pobranie danych z repozytorium
	devicesList, count, err := repository.GetDevices(db.DB, singular, filters, page)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	// Meta dla paginacji
	const perPage = 20
	meta := &apiresponse.Meta{
		Count:       count,
		CurrentPage: page,
		TotalPages:  (count + perPage - 1) / perPage,
		PerPage:     perPage,
	}

	apiresponse.JSONSuccess(w, http.StatusOK, devicesList, meta)
}