package repository

import (
	"database/sql"
	"time"

	"github.com/superstan777/stock-backend/internal/tickets"
)

// GetAll pobiera wszystkie tickety.
func GetAll(db *sql.DB) ([]tickets.Ticket, error) {
	rows, err := db.Query(`
		SELECT id, number, title, description, caller_id, assigned_to, status, created_at, estimated_resolution_date, resolution_date
		FROM tickets
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []tickets.Ticket
	for rows.Next() {
		var t tickets.Ticket
		if err := rows.Scan(
			&t.ID, &t.Number, &t.Title, &t.Description,
			&t.CallerID, &t.AssignedTo, &t.Status,
			&t.CreatedAt, &t.EstimatedResolutionDate, &t.ResolutionDate,
		); err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, nil
}

// GetByID zwraca jeden ticket po ID.
func GetByID(db *sql.DB, id string) (*tickets.Ticket, error) {
	row := db.QueryRow(`
		SELECT id, number, title, description, caller_id, assigned_to, status, created_at, estimated_resolution_date, resolution_date
		FROM tickets
		WHERE id = $1
	`, id)

	var t tickets.Ticket
	if err := row.Scan(
		&t.ID, &t.Number, &t.Title, &t.Description,
		&t.CallerID, &t.AssignedTo, &t.Status,
		&t.CreatedAt, &t.EstimatedResolutionDate, &t.ResolutionDate,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

// Insert dodaje nowy ticket.
func Insert(db *sql.DB, input tickets.TicketInsert) (*tickets.Ticket, error) {
	createdAt := time.Now()
	if input.CreatedAt != nil {
		createdAt = *input.CreatedAt
	}

	row := db.QueryRow(`
		INSERT INTO tickets (title, description, caller_id, assigned_to, status, created_at, estimated_resolution_date, resolution_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, number, title, description, caller_id, assigned_to, status, created_at, estimated_resolution_date, resolution_date
	`,
		input.Title, input.Description, input.CallerID, input.AssignedTo, input.Status, createdAt, input.EstimatedResolutionDate, input.ResolutionDate,
	)

	var t tickets.Ticket
	if err := row.Scan(
		&t.ID, &t.Number, &t.Title, &t.Description,
		&t.CallerID, &t.AssignedTo, &t.Status,
		&t.CreatedAt, &t.EstimatedResolutionDate, &t.ResolutionDate,
	); err != nil {
		return nil, err
	}
	return &t, nil
}

// Update aktualizuje ticket.
func Update(db *sql.DB, id string, input tickets.TicketUpdate) (*tickets.Ticket, error) {
	row := db.QueryRow(`
		UPDATE tickets
		SET title = COALESCE($1, title),
		    description = COALESCE($2, description),
		    caller_id = COALESCE($3, caller_id),
		    assigned_to = COALESCE($4, assigned_to),
		    status = COALESCE($5, status),
		    estimated_resolution_date = COALESCE($6, estimated_resolution_date),
		    resolution_date = COALESCE($7, resolution_date)
		WHERE id = $8
		RETURNING id, number, title, description, caller_id, assigned_to, status, created_at, estimated_resolution_date, resolution_date
	`,
		input.Title, input.Description, input.CallerID, input.AssignedTo, input.Status, input.EstimatedResolutionDate, input.ResolutionDate, id,
	)

	var t tickets.Ticket
	if err := row.Scan(
		&t.ID, &t.Number, &t.Title, &t.Description,
		&t.CallerID, &t.AssignedTo, &t.Status,
		&t.CreatedAt, &t.EstimatedResolutionDate, &t.ResolutionDate,
	); err != nil {
		return nil, err
	}
	return &t, nil
}

// Delete usuwa ticket po ID.
func Delete(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM tickets WHERE id = $1`, id)
	return err
}