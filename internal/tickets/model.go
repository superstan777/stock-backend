package tickets

import "time"

// Ticket odpowiada strukturze wiersza w tabeli tickets
type Ticket struct {
	ID                     string     `db:"id"`
	Number                 int        `db:"number"`
	Title                  string     `db:"title"`
	Description            *string    `db:"description"`
	CallerID               *string    `db:"caller_id"`
	AssignedTo             *string    `db:"assigned_to"`
	Status                 string     `db:"status"`
	CreatedAt              time.Time  `db:"created_at"`
	EstimatedResolutionDate *time.Time `db:"estimated_resolution_date"`
	ResolutionDate         *time.Time `db:"resolution_date"`
}

// TicketInsert używany przy tworzeniu nowego ticketa
type TicketInsert struct {
	Number                 *int       `db:"number,omitempty"`
	Title                  string     `db:"title"`
	Description            *string    `db:"description,omitempty"`
	CallerID               *string    `db:"caller_id,omitempty"`
	AssignedTo             *string    `db:"assigned_to,omitempty"`
	Status                 string     `db:"status"`
	CreatedAt              *time.Time `db:"created_at,omitempty"`
	EstimatedResolutionDate *time.Time `db:"estimated_resolution_date,omitempty"`
	ResolutionDate         *time.Time `db:"resolution_date,omitempty"`
}

// TicketUpdate używany przy aktualizacji ticketa
type TicketUpdate struct {
	Title                  *string    `db:"title,omitempty"`
	Description            *string    `db:"description,omitempty"`
	CallerID               *string    `db:"caller_id,omitempty"`
	AssignedTo             *string    `db:"assigned_to,omitempty"`
	Status                 *string    `db:"status,omitempty"`
	EstimatedResolutionDate *time.Time `db:"estimated_resolution_date,omitempty"`
	ResolutionDate         *time.Time `db:"resolution_date,omitempty"`
}