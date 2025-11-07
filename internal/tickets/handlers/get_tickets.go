package handlers

import (
	"net/http"
	"strconv"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/tickets/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	page := 1
	if p := query.Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	perPage := 20
	if p := query.Get("per_page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			perPage = parsed
		}
	}

	filters := make(map[string]string)
	for key, values := range query {
		if key == "page" || key == "per_page" {
			continue
		}
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}

	ticketsList, total, err := repository.GetTickets(db.DB, filters, page, perPage)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	meta := map[string]interface{}{
		"count":        total,
		"current_page": page,
		"total_pages":  (total + perPage - 1) / perPage, 
		"per_page":     perPage,
	}

	apiresponse.JSONSuccess(w, http.StatusOK, "Tickets fetched successfully", map[string]interface{}{
		"tickets": ticketsList,
		"meta":    meta,
	})
}