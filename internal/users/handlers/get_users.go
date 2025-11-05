package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/users/repository"
)

// GetUsersHandler obsługuje GET /api/users z filtrami i paginacją.
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// --- PAGINACJA ---
	page := 1
	perPage := 20

	if p := r.URL.Query().Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if pp := r.URL.Query().Get("perPage"); pp != "" {
		if parsed, err := strconv.Atoi(pp); err == nil && parsed > 0 {
			perPage = parsed
		}
	}

	// --- FILTRY ---
	filters := make(map[string]string)
	for key, values := range r.URL.Query() {
		// Pomijamy page i perPage
		if key == "page" || key == "perPage" {
			continue
		}
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	// --- POBRANIE DANYCH Z REPO ---
	usersList, count, err := repository.GetUsers(db.DB, filters, page, perPage)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// --- ODPOWIEDŹ ---
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  usersList,
		"count": count,
	})
}