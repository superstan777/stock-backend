package handlers

import (
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	statsRepo "github.com/superstan777/stock-backend/internal/tickets/stats/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func GetResolvedTicketsStatsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := statsRepo.GetResolvedTicketsStats(db.DB)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	// Zwracamy dane statystyczne, meta nil bo brak paginacji
	apiresponse.JSONSuccess(w, http.StatusOK, data, nil)
}