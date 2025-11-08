package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/superstan777/stock-backend/internal/tickets"
)

// ==============================
// CRUD + STATYSTYKI TICKETÓW
// ==============================

// GetTickets – pobiera tickety z filtrowaniem, paginacją i JOINami do userów
func GetTickets(db *sql.DB, filters map[string]string, page, perPage int) ([]tickets.TicketWithUsers, int, error) {
	baseQuery := `
		SELECT 
			t.id,
			t.number,
			t.title,
			t.description,
			t.status,
			t.created_at,
			t.estimated_resolution_date,
			t.resolution_date,
			c.id AS caller_id,
			c.email AS caller_email,
			a.id AS assigned_id,
			a.email AS assigned_email
		FROM tickets t
		LEFT JOIN users c ON t.caller_id = c.id
		LEFT JOIN users a ON t.assigned_to = a.id
		WHERE 1=1
	`

	args := []interface{}{}
	argIdx := 1

	for key, value := range filters {
		if value == "" {
			continue
		}

		val := strings.TrimSpace(value)
		switch key {
		case "status":
			// --- obsługa wielu statusów jak w GetDevices ---
			statuses := strings.Split(val, ",")
			placeholders := []string{}
			for _, s := range statuses {
				s = strings.TrimSpace(s)
				placeholders = append(placeholders, fmt.Sprintf("$%d::ticket_status", argIdx))
				args = append(args, s)
				argIdx++
			}
			baseQuery += fmt.Sprintf(" AND t.status IN (%s)", strings.Join(placeholders, ","))
		case "number":
			baseQuery += fmt.Sprintf(" AND t.number = $%d", argIdx)
			args = append(args, val)
			argIdx++
		case "title":
			baseQuery += fmt.Sprintf(" AND t.title ILIKE $%d", argIdx)
			args = append(args, "%"+val+"%")
			argIdx++
		case "caller_email":
			baseQuery += fmt.Sprintf(" AND c.email ILIKE $%d", argIdx)
			args = append(args, "%"+val+"%")
			argIdx++
		case "assigned_email":
			baseQuery += fmt.Sprintf(" AND a.email ILIKE $%d", argIdx)
			args = append(args, "%"+val+"%")
			argIdx++
		}
	}

	offset := (page - 1) * perPage
	baseQuery += fmt.Sprintf(" ORDER BY t.created_at DESC LIMIT %d OFFSET %d", perPage, offset)

	rows, err := db.Query(baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var ticketsList []tickets.TicketWithUsers
	for rows.Next() {
		var t tickets.TicketWithUsers
		var callerID, assignedID sql.NullString
		var callerEmail, assignedEmail sql.NullString

		if err := rows.Scan(
			&t.ID, &t.Number, &t.Title, &t.Description, &t.Status,
			&t.CreatedAt, &t.EstimatedResolutionDate, &t.ResolutionDate,
			&callerID, &callerEmail, &assignedID, &assignedEmail,
		); err != nil {
			return nil, 0, err
		}

		if callerID.Valid {
			t.Caller = &tickets.User{ID: callerID.String, Email: callerEmail.String}
		}
		if assignedID.Valid {
			t.AssignedTo = &tickets.User{ID: assignedID.String, Email: assignedEmail.String}
		}
		ticketsList = append(ticketsList, t)
	}

	// --- liczenie wszystkich pasujących rekordów ---
	countQuery := "SELECT COUNT(*) FROM tickets t WHERE 1=1"
	countArgs := []interface{}{}
	argIdx = 1

	for key, value := range filters {
		if value == "" {
			continue
		}
		val := strings.TrimSpace(value)
		switch key {
		case "status":
			statuses := strings.Split(val, ",")
			placeholders := []string{}
			for _, s := range statuses {
				s = strings.TrimSpace(s)
				placeholders = append(placeholders, fmt.Sprintf("$%d::ticket_status", argIdx))
				countArgs = append(countArgs, s)
				argIdx++
			}
			countQuery += fmt.Sprintf(" AND t.status IN (%s)", strings.Join(placeholders, ","))
		case "number":
			countQuery += fmt.Sprintf(" AND t.number = $%d", argIdx)
			countArgs = append(countArgs, val)
			argIdx++
		case "title":
			countQuery += fmt.Sprintf(" AND t.title ILIKE $%d", argIdx)
			countArgs = append(countArgs, "%"+val+"%")
			argIdx++
		case "caller_email":
			countQuery += fmt.Sprintf(" AND c.email ILIKE $%d", argIdx)
			countArgs = append(countArgs, "%"+val+"%")
			argIdx++
		case "assigned_email":
			countQuery += fmt.Sprintf(" AND a.email ILIKE $%d", argIdx)
			countArgs = append(countArgs, "%"+val+"%")
			argIdx++
		}
	}

	var totalCount int
	if err := db.QueryRow(countQuery, countArgs...).Scan(&totalCount); err != nil {
		return nil, 0, err
	}

	return ticketsList, totalCount, nil
}

// ----------------------
// CRUD (GetByID, Insert, Update, Delete)
// ----------------------

func GetByID(db *sql.DB, id string) (*tickets.TicketWithUsers, error) {
	query := `
		SELECT 
			t.id,
			t.number,
			t.title,
			t.description,
			t.status,
			t.created_at,
			t.estimated_resolution_date,
			t.resolution_date,
			c.id AS caller_id,
			c.email AS caller_email,
			a.id AS assigned_id,
			a.email AS assigned_email
		FROM tickets t
		LEFT JOIN users c ON t.caller_id = c.id
		LEFT JOIN users a ON t.assigned_to = a.id
		WHERE t.id = $1
	`

	row := db.QueryRow(query, id)

	var t tickets.TicketWithUsers
	var callerID, assignedID sql.NullString
	var callerEmail, assignedEmail sql.NullString

	if err := row.Scan(
		&t.ID, &t.Number, &t.Title, &t.Description, &t.Status,
		&t.CreatedAt, &t.EstimatedResolutionDate, &t.ResolutionDate,
		&callerID, &callerEmail, &assignedID, &assignedEmail,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if callerID.Valid {
		t.Caller = &tickets.User{ID: callerID.String, Email: callerEmail.String}
	}
	if assignedID.Valid {
		t.AssignedTo = &tickets.User{ID: assignedID.String, Email: assignedEmail.String}
	}

	return &t, nil
}

func Insert(db *sql.DB, input tickets.TicketInsert) (*tickets.Ticket, error) {
	row := db.QueryRow(`
		INSERT INTO tickets (title, description, caller_id)
		VALUES ($1, $2, $3)
		RETURNING id, number, title, description, caller_id, assigned_to, status, created_at, estimated_resolution_date, resolution_date
	`,
		input.Title, input.Description, input.CallerID,
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

func Update(db *sql.DB, id string, input tickets.TicketUpdate) (*tickets.Ticket, error) {
	row := db.QueryRow(`
		UPDATE tickets SET
		    title = COALESCE($1, title),
		    description = COALESCE($2, description),
		    caller_id = COALESCE($3, caller_id),
		    assigned_to = COALESCE($4, assigned_to),
		    status = COALESCE($5, status),
		    estimated_resolution_date = COALESCE($6, estimated_resolution_date),
		    resolution_date = COALESCE($7, resolution_date)
		WHERE id = $8
		RETURNING id, number, title, description, caller_id, assigned_to, status, created_at, estimated_resolution_date, resolution_date
	`,
		input.Title, input.Description, input.CallerID, input.AssignedTo,
		input.Status, input.EstimatedResolutionDate, input.ResolutionDate, id,
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

func Delete(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM tickets WHERE id = $1`, id)
	return err
}
