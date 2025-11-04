package users

import "time"

// User reprezentuje rekord w tabeli "users".
type User struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
}

// UserInput reprezentuje dane wejściowe przy tworzeniu użytkownika.
type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserUpdate reprezentuje dane wejściowe przy aktualizacji użytkownika.
type UserUpdate struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}