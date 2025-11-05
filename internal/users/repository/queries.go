package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/superstan777/stock-backend/internal/users"
)

// GetUsers pobiera użytkowników z filtrami i paginacją.
func GetUsers(db *sql.DB, filters map[string]string, page, perPage int) ([]users.User, int, error) {
	baseQuery := `SELECT id, name, email, created_at FROM users`
	countQuery := `SELECT COUNT(*) FROM users`

	// ---- FILTRY ----
	var whereClauses []string
	var args []interface{}
	argIndex := 1

	for key, value := range filters {
		if value == "" {
			continue
		}

		values := strings.Split(value, ",")
		for i := range values {
			values[i] = strings.TrimSpace(values[i])
		}

		if len(values) > 1 {
			var ors []string
			for _, v := range values {
				ors = append(ors, fmt.Sprintf("%s ILIKE $%d", key, argIndex))
				args = append(args, v+"%")
				argIndex++
			}
			whereClauses = append(whereClauses, "("+strings.Join(ors, " OR ")+")")
		} else {
			whereClauses = append(whereClauses, fmt.Sprintf("%s ILIKE $%d", key, argIndex))
			args = append(args, values[0]+"%")
			argIndex++
		}
	}

	if len(whereClauses) > 0 {
		where := " WHERE " + strings.Join(whereClauses, " AND ")
		baseQuery += where
		countQuery += where
	}

	// ---- SORTOWANIE ----
	baseQuery += " ORDER BY name ASC"

	// ---- PAGINACJA ----
	offset := (page - 1) * perPage
	baseQuery += fmt.Sprintf(" OFFSET %d LIMIT %d", offset, perPage)

	// ---- LICZENIE ----
	var total int
	if err := db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// ---- POBRANIE DANYCH ----
	rows, err := db.Query(baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []users.User
	for rows.Next() {
		var u users.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, 0, err
		}
		list = append(list, u)
	}

	return list, total, nil
}


// GetByID zwraca jednego użytkownika po ID.
func GetByID(db *sql.DB, id string) (*users.User, error) {
	row := db.QueryRow(`SELECT id, name, email, created_at FROM users WHERE id = $1`, id)

	var u users.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

// Insert dodaje nowego użytkownika.
func Insert(db *sql.DB, input users.UserInput) (*users.User, error) {
	row := db.QueryRow(
		`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at`,
		input.Name, input.Email,
	)
	var u users.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// Update aktualizuje użytkownika.
func Update(db *sql.DB, id string, input users.UserUpdate) (*users.User, error) {
	row := db.QueryRow(
		`UPDATE users 
		 SET name = COALESCE($1, name),
		     email = COALESCE($2, email)
		 WHERE id = $3
		 RETURNING id, name, email, created_at`,
		input.Name, input.Email, id,
	)
	var u users.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// Delete usuwa użytkownika po ID.
func Delete(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}