package handlers

import (
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	statsRepo "github.com/superstan777/stock-backend/internal/tickets/stats/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func GetTicketsByOperatorHandler(w http.ResponseWriter, r *http.Request) {
	data, err := statsRepo.GetTicketsByOperator(db.DB)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	apiresponse.JSONSuccess(w, http.StatusOK, data, nil)
}