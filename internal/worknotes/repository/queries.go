package repository

import (
	"database/sql"
	"time"

	"github.com/superstan777/stock-backend/internal/worknotes"
)


func GetWorknotesByTicket(db *sql.DB, ticketID string) ([]worknotes.WorknoteWithAuthor, error) {
	query := `
	SELECT 
		w.id,
		w.ticket_id,
		w.note,
		w.created_at,
		u.id AS author_id,
		u.email AS author_email
	FROM worknotes w
	LEFT JOIN users u ON w.author_id = u.id
	WHERE w.ticket_id = $1
	ORDER BY w.created_at DESC;
	`

	rows, err := db.Query(query, ticketID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []worknotes.WorknoteWithAuthor

	for rows.Next() {
		var n worknotes.WorknoteWithAuthor
		var a worknotes.Author
		if err := rows.Scan(&n.ID, &n.TicketID, &n.Note, &n.CreatedAt, &a.ID, &a.Email); err != nil {
			return nil, err
		}
		n.Author = a
		notes = append(notes, n)
	}

	return notes, nil
}

func AddWorknote(db *sql.DB, note worknotes.WorknoteInsert) (*worknotes.Worknote, error) {
	query := `
	INSERT INTO worknotes (ticket_id, author_id, note, created_at)
	VALUES ($1, $2, $3, $4)
	RETURNING id, ticket_id, author_id, note, created_at;
	`

	now := time.Now().UTC()
	row := db.QueryRow(query, note.TicketID, note.AuthorID, note.Note, now)

	var newNote worknotes.Worknote
	if err := row.Scan(&newNote.ID, &newNote.TicketID, &newNote.AuthorID, &newNote.Note, &newNote.CreatedAt); err != nil {
		return nil, err
	}

	return &newNote, nil
}