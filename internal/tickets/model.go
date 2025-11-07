package tickets

import "time"

type Ticket struct {
	ID                      string     `db:"id" json:"id"`
	Number                  int        `db:"number" json:"number"`
	Title                   string     `db:"title" json:"title"`
	Description              *string    `db:"description" json:"description"`
	CallerID                *string    `db:"caller_id" json:"caller_id"`
	AssignedTo              *string    `db:"assigned_to" json:"assigned_to"`
	Status                  string     `db:"status" json:"status"`
	CreatedAt               time.Time  `db:"created_at" json:"created_at"`
	EstimatedResolutionDate *time.Time `db:"estimated_resolution_date" json:"estimated_resolution_date"`
	ResolutionDate          *time.Time `db:"resolution_date" json:"resolution_date"`
}

type TicketInsert struct {
	Title       string  `db:"title" json:"title"`
	Description *string `db:"description" json:"description"`
	CallerID    *string `db:"caller_id" json:"caller_id"`
}

type TicketUpdate struct {
	Title                   *string    `db:"title" json:"title"`
	Description             *string    `db:"description" json:"description"`
	CallerID                *string    `db:"caller_id" json:"caller_id"`
	AssignedTo              *string    `db:"assigned_to" json:"assigned_to"`
	Status                  *string    `db:"status" json:"status"`
	EstimatedResolutionDate *time.Time `db:"estimated_resolution_date" json:"estimated_resolution_date"`
	ResolutionDate          *time.Time `db:"resolution_date" json:"resolution_date"`
}

type User struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Email string  `json:"email"`
}

type TicketWithUsers struct {
	ID                      string     `json:"id"`
	Number                  int        `json:"number"`
	Title                   string     `json:"title"`
	Description             *string    `json:"description"`
	Status                  string     `json:"status"`
	CreatedAt               time.Time  `json:"created_at"`
	EstimatedResolutionDate *time.Time `json:"estimated_resolution_date"`
	ResolutionDate          *time.Time `json:"resolution_date"`
	Caller                  *User      `json:"caller"`
	AssignedTo              *User      `json:"assigned_to"`
}