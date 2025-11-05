package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/superstan777/stock-backend/internal/db"
	"github.com/superstan777/stock-backend/internal/worknotes"
	"github.com/superstan777/stock-backend/internal/worknotes/repository"
)

// POST /worknotes
func CreateWorknoteHandler(w http.ResponseWriter, r *http.Request) {
	var note worknotes.WorknoteInsert
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if note.TicketID == "" || note.AuthorID == "" || note.Note == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	newNote, err := repository.AddWorknote(db.DB, note)
	if err != nil {
		http.Error(w, "DB insert error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newNote)
}