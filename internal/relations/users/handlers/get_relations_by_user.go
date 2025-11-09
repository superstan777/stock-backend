package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/relations/users/repository"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
)

func GetRelationsByUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	if userID == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing user_id")
		return
	}

	list, err := repository.GetByUser(db.DB, userID)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database query error: "+err.Error())
		return
	}

	// Zwracamy dane bez komunikatu, meta nil bo brak paginacji
	apiresponse.JSONSuccess(w, http.StatusOK, list, nil)
}