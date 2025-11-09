package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/relations/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func EndRelationHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing relation ID")
		return
	}

	if err := repository.End(db.DB, id); err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database update error: "+err.Error())
		return
	}

	// Użycie nowego helpera JSONEnded
	apiresponse.JSONEnded(w, id)
}