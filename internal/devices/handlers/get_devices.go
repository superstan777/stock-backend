package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/devices/repository"
)

// GetDevicesHandler obsługuje GET /api/devices/{device_type}
// Przykład: /api/devices/computers?page=1&serial_number=ABC
func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	deviceType := chi.URLParam(r, "device_type")

	query := r.URL.Query()

	// --- jeśli deviceType jest pusty, to zwracamy wszystkie ---
	if deviceType == "" {
		// pobierz wszystkie urządzenia (np. bez filtra device_type)
		devicesList, count, err := repository.GetDevices(db.DB, "", nil, 1)
		if err != nil {
			http.Error(w, "DB query error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data":  devicesList,
			"count": count,
		})
		return
	}

	// --- walidacja liczby mnogiej ---
	validTypes := map[string]string{
		"computers": "computer",
		"monitors":  "monitor",
	}

	singular, ok := validTypes[deviceType]
	if !ok {
		http.Error(w, "Invalid device type", http.StatusBadRequest)
		return
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

	// --- PAGINACJA ---
	page := 1
	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	// --- POBRANIE DANYCH ---
	devicesList, count, err := repository.GetDevices(db.DB, singular, filters, page)
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