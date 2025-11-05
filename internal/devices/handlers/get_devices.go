package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
)

// GetDevicesHandler obsługuje GET /api/devices
// Można filtrować query params, np: ?device_type=Computer&serial_number=ABC&page=1
func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	deviceType := query.Get("device_type") // opcjonalny

	// --- FILTRY ---
	filters := make(map[string]string)
	for key, values := range query {
		if key == "page" || key == "device_type" {
			continue
		}
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	// --- PAGINACJA ---
	page := 1
	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	// --- POBRANIE DANYCH ---
	devicesList, count, err := repository.GetDevices(db.DB, deviceType, filters, page)
	if err != nil {
		http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// --- ODPOWIEDŹ JSON ---
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  devicesList,
		"count": count,
	})
}