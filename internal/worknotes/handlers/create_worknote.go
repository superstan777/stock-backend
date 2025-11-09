package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/utils/apiresponse"
	"github.com/superstan777/stock-backend/internal/worknotes"
	"github.com/superstan777/stock-backend/internal/worknotes/repository"
)

func CreateWorknoteHandler(w http.ResponseWriter, r *http.Request) {
	var note worknotes.WorknoteInsert
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		apiresponse.JSONError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if note.TicketID == "" || note.AuthorID == "" || note.Note == "" {
		apiresponse.JSONError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	newNote, err := repository.AddWorknote(db.DB, note)
	if err != nil {
		apiresponse.JSONError(w, http.StatusInternalServerError, "Database insert error: "+err.Error())
		return
	}

	apiresponse.JSONCreated(w, newNote)
}