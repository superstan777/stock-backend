package worknotes

import "time"

type Worknote struct {
	ID        string    `json:"id" db:"id"`
	TicketID  string    `json:"ticket_id" db:"ticket_id"`
	AuthorID  string    `json:"author_id" db:"author_id"`
	Note      string    `json:"note" db:"note"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Author struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type WorknoteWithAuthor struct {
	ID        string    `json:"id"`
	TicketID  string    `json:"ticket_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	Author    Author    `json:"author"`
}

type WorknoteInsert struct {
	TicketID string `json:"ticket_id"`
	AuthorID string `json:"author_id"`
	Note     string `json:"note"`
}