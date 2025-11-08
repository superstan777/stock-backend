package repository

import (
	"database/sql"
	"time"

	"github.com/superstan777/stock-backend/internal/relations"
)

func Insert(db *sql.DB, input relations.RelationInsert) (*relations.Relation, error) {
	query := `
		INSERT INTO relations (device_id, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id, device_id, user_id, start_date, end_date
	`

	row := db.QueryRow(query, input.DeviceID, input.UserID, input.StartDate, input.EndDate)
	var rel relations.Relation
	if err := row.Scan(&rel.ID, &rel.DeviceID, &rel.UserID, &rel.StartDate, &rel.EndDate); err != nil {
		return nil, err
	}
	return &rel, nil
}

func End(db *sql.DB, id string) error {
	query := `UPDATE relations SET end_date = $1 WHERE id = $2`
	_, err := db.Exec(query, time.Now(), id)
	return err
}

