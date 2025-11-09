package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

// GetDevicesHandler obsługuje GET /api/devices/{device_type}
// Przykład: /api/devices/computers?page=1&serial_number=ABC
func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	deviceType := chi.URLParam(r, "device_type")
	query := r.URL.Query()

	// --- Mapowanie liczby mnogiej na pojedynczą ---
	validTypes := map[string]string{
		"computers": "computer",
		"monitors":  "monitor",
	}

	// --- Parsowanie parametru ?page= ---
	page := 1
	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	// --- FILTRY ---
	filters := make(map[string]string)
	for key, values := range query {
		if key == "page" {
			continue
		}
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	// --- Walidacja device_type ---
	var singular string
	if deviceType != "" {
		var ok bool
		singular, ok = validTypes[deviceType]
		if !ok {
			apiresponse.JSONError(w, http.StatusBadRequest, "Invalid device type: "+deviceType)
			return
		}
	}

	// --- Pobranie danych z repozytorium ---
	devicesList, count, err := repository.GetDevices(db.DB, singular, filters, page)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	// --- Obliczanie metadanych paginacji ---
	meta := &apiresponse.Meta{
		Count:       count,
		CurrentPage: page,
		TotalPages:  (count + 20 - 1) / 20,
		PerPage:     20,
	}

	// --- Zwrócenie ujednoliconej odpowiedzi ---
	apiresponse.JSONSuccess(w, http.StatusOK, devicesList, meta)
}