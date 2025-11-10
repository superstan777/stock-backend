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
			o.id AS operator_id,
			o.email AS operator_email
		FROM tickets t
		LEFT JOIN users c ON t.caller_id = c.id
		LEFT JOIN users o ON t.operator_id = o.id
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
		case "operator_email":
			baseQuery += fmt.Sprintf(" AND o.email ILIKE $%d", argIdx)
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
		var callerID, operatorID sql.NullString
		var callerEmail, operatorEmail sql.NullString

		if err := rows.Scan(
			&t.ID, &t.Number, &t.Title, &t.Description, &t.Status,
			&t.CreatedAt, &t.EstimatedResolutionDate, &t.ResolutionDate,
			&callerID, &callerEmail, &operatorID, &operatorEmail,
		); err != nil {
			return nil, 0, err
		}

		if callerID.Valid {
			t.Caller = &tickets.User{ID: callerID.String, Email: callerEmail.String}
		}
		if operatorID.Valid {
			t.Operator = &tickets.User{ID: operatorID.String, Email: operatorEmail.String}
		}
		ticketsList = append(ticketsList, t)
	}

	// --- liczenie wszystkich pasujących rekordów ---
	countQuery := `
		SELECT COUNT(*) 
		FROM tickets t
		LEFT JOIN users c ON t.caller_id = c.id
		LEFT JOIN users o ON t.operator_id = o.id
		WHERE 1=1
	`
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
		case "operator_email":
			countQuery += fmt.Sprintf(" AND o.email ILIKE $%d", argIdx)
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
			o.id AS operator_id,
			o.email AS operator_email
		FROM tickets t
		LEFT JOIN users c ON t.caller_id = c.id
		LEFT JOIN users o ON t.operator_id = o.id
		WHERE t.id = $1
	`

	row := db.QueryRow(query, id)

	var t tickets.TicketWithUsers
	var callerID, operatorID sql.NullString
	var callerEmail, operatorEmail sql.NullString

	if err := row.Scan(
		&t.ID, &t.Number, &t.Title, &t.Description, &t.Status,
		&t.CreatedAt, &t.EstimatedResolutionDate, &t.ResolutionDate,
		&callerID, &callerEmail, &operatorID, &operatorEmail,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if callerID.Valid {
		t.Caller = &tickets.User{ID: callerID.String, Email: callerEmail.String}
	}
	if operatorID.Valid {
		t.Operator = &tickets.User{ID: operatorID.String, Email: operatorEmail.String}
	}

	return &t, nil
}

func Insert(db *sql.DB, input tickets.TicketInsert) (*tickets.Ticket, error) {
	row := db.QueryRow(`
		INSERT INTO tickets (title, description, caller_id)
		VALUES ($1, $2, $3)
		RETURNING id, number, title, description, caller_id, operator_id, status, created_at, estimated_resolution_date, resolution_date
	`,
		input.Title, input.Description, input.CallerID,
	)

	var t tickets.Ticket
	if err := row.Scan(
		&t.ID, &t.Number, &t.Title, &t.Description,
		&t.CallerID, &t.OperatorID, &t.Status,
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
		    operator_id = COALESCE($4, operator_id),
		    status = COALESCE($5, status),
		    estimated_resolution_date = COALESCE($6, estimated_resolution_date),
		    resolution_date = COALESCE($7, resolution_date)
		WHERE id = $8
		RETURNING id, number, title, description, caller_id, operator_id, status, created_at, estimated_resolution_date, resolution_date
	`,
		input.Title, input.Description, input.CallerID, input.OperatorID,
		input.Status, input.EstimatedResolutionDate, input.ResolutionDate, id,
	)

	var t tickets.Ticket
	if err := row.Scan(
		&t.ID, &t.Number, &t.Title, &t.Description,
		&t.CallerID, &t.OperatorID, &t.Status,
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