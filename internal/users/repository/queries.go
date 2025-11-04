package repository

import (
	"database/sql"

	"github.com/superstan777/stock-backend/internal/users"
)

// GetAll pobiera wszystkich użytkowników.
func GetAll(db *sql.DB) ([]users.User, error) {
	rows, err := db.Query(`SELECT id, name, email, created_at FROM users ORDER BY name ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []users.User
	for rows.Next() {
		var u users.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		list = append(list, u)
	}
	return list, nil
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