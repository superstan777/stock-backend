package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
)

// GetDevicesHandler obsługuje GET /api/devices
// Można filtrować query params, np: ?device_type=Computer&serial_number=ABC&page=1&perPage=20
func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	deviceType := query.Get("device_type") // teraz opcjonalny

	// --- filtrujemy wszystkie query params oprócz page, perPage i device_type ---
	filters := make(map[string]string)
	for key, values := range query {
		if key == "page" || key == "perPage" || key == "device_type" {
			continue
		}
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	// --- paginacja ---
	page := 1
	perPage := 20

	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	if pp := query.Get("perPage"); pp != "" {
		if parsed, err := strconv.Atoi(pp); err == nil && parsed > 0 {
			perPage = parsed
		}
	}

	// --- pobieranie danych z repozytorium ---
	devicesList, count, err := repository.GetDevices(db.DB, deviceType, filters, page, perPage)
	if err != nil {
		http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// --- odpowiedź JSON ---
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  devicesList,
		"count": count,
	})
}