package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/tickets/repository"
)

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	filters := map[string]string{}
	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage, _ := strconv.Atoi(r.URL.Query().Get("per_page"))
	if perPage <= 0 {
		perPage = 20
	}

	list, total, err := repository.GetTickets(db.DB, filters, page, perPage)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"data":  list,
		"total": total,
		"page":  page,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}