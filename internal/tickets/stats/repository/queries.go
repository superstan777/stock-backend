package repository

import (
	"database/sql"

	"github.com/superstan777/stock-backend/internal/tickets/stats"
)



func GetResolvedTicketsStats(db *sql.DB) ([]stats.ResolvedTicketsStats, error) {
	query := `
		SELECT 
			TO_CHAR(DATE(t.resolution_date AT TIME ZONE 'UTC'), 'YYYY-MM-DD') AS date,
			COUNT(*) AS count
		FROM tickets t
		WHERE t.resolution_date IS NOT NULL
		  AND t.status = 'resolved'
		GROUP BY DATE(t.resolution_date AT TIME ZONE 'UTC')
		ORDER BY DATE(t.resolution_date AT TIME ZONE 'UTC') ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []stats.ResolvedTicketsStats
	for rows.Next() {
		var s stats.ResolvedTicketsStats
		if err := rows.Scan(&s.Date, &s.Count); err != nil {
			return nil, err
		}
		data = append(data, s)
	}
	return data, nil
}

func GetOpenTicketsStats(db *sql.DB) ([]stats.OpenTicketsStats, error) {
	query := `
		SELECT 
			CASE 
				WHEN t.estimated_resolution_date IS NULL THEN 'No ETA'
				ELSE TO_CHAR(DATE(t.estimated_resolution_date AT TIME ZONE 'UTC'), 'YYYY-MM-DD')
			END AS date,
			COUNT(*) AS count
		FROM tickets t
		WHERE t.status IN ('new', 'on_hold', 'in_progress')
		GROUP BY 1
		ORDER BY MIN(
			CASE 
				WHEN t.estimated_resolution_date IS NULL THEN '9999-12-31' 
				ELSE TO_CHAR(DATE(t.estimated_resolution_date AT TIME ZONE 'UTC'), 'YYYY-MM-DD')
			END
		)
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []stats.OpenTicketsStats
	for rows.Next() {
		var s stats.OpenTicketsStats
		if err := rows.Scan(&s.Date, &s.Count); err != nil {
			return nil, err
		}
		data = append(data, s)
	}
	return data, nil
}

func GetTicketsByOperator(db *sql.DB) ([]stats.OperatorTicketsStats, error) {
	query := `
		SELECT 
			u.id,
			u.name,
			u.email,
			COUNT(*) AS count
		FROM tickets t
		LEFT JOIN users u ON t.operator_id = u.id
		WHERE t.status NOT IN ('resolved', 'cancelled')
		GROUP BY u.id, u.name, u.email
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []stats.OperatorTicketsStats
	for rows.Next() {
		var s stats.OperatorTicketsStats
		var id, name, email sql.NullString

		if err := rows.Scan(&id, &name, &email, &s.Count); err != nil {
			return nil, err
		}

		if id.Valid {
			s.Operator.ID = &id.String
		} else {
			s.Operator.ID = nil
		}

		if name.Valid {
			s.Operator.Name = &name.String
		} else {
			s.Operator.Name = nil
		}

		if email.Valid {
			s.Operator.Email = &email.String
		} else {
			s.Operator.Email = nil
		}

		data = append(data, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}