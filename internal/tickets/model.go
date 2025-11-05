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


type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type TicketWithUsers struct {
	ID                      string     `json:"id"`
	Number                  int        `json:"number"`
	Title                   string     `json:"title"`
	Description              *string    `json:"description,omitempty"`
	Status                  string     `json:"status"`
	CreatedAt               time.Time  `json:"created_at"`
	EstimatedResolutionDate *time.Time `json:"estimated_resolution_date,omitempty"`
	ResolutionDate          *time.Time `json:"resolution_date,omitempty"`
	Caller                  *User      `json:"caller,omitempty"`
	AssignedTo              *User      `json:"assigned_to,omitempty"`
}