package handlers

import (
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	statsRepo "github.com/superstan777/stock-backend/internal/tickets/stats/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func GetOpenTicketsStatsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := statsRepo.GetOpenTicketsStats(db.DB)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	apiresponse.JSONSuccess(w, http.StatusOK, "Open tickets stats fetched successfully", data)
}