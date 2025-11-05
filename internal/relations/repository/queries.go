package repository

import (
	"database/sql"
	"time"

	"github.com/superstan777/stock-backend/internal/relations"
)

// --- GET BY DEVICE ---
func GetByDevice(db *sql.DB, deviceID string) ([]relations.RelationWithDetails, error) {
	query := `
		SELECT 
			r.id,
			r.start_date,
			r.end_date,
			u.id AS user_id,
			u.email,
			u.name,
			d.id AS device_id,
			d.model,
			d.serial_number,
			d.device_type,
			d.install_status
		FROM relations r
		LEFT JOIN users u ON r.user_id = u.id
		LEFT JOIN devices d ON r.device_id = d.id
		WHERE r.device_id = $1
		ORDER BY r.start_date DESC
	`

	rows, err := db.Query(query, deviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []relations.RelationWithDetails
	for rows.Next() {
		var rel relations.RelationWithDetails
		if err := rows.Scan(
			&rel.ID,
			&rel.StartDate,
			&rel.EndDate,
			&rel.User.ID,
			&rel.User.Email,
			&rel.User.Name,
			&rel.Device.ID,
			&rel.Device.Model,
			&rel.Device.SerialNumber,
			&rel.Device.DeviceType,
			&rel.Device.InstallStatus,
		); err != nil {
			return nil, err
		}
		list = append(list, rel)
	}
	return list, nil
}

// --- GET BY USER ---
func GetByUser(db *sql.DB, userID string) ([]relations.RelationWithDetails, error) {
	query := `
		SELECT 
			r.id,
			r.start_date,
			r.end_date,
			u.id AS user_id,
			u.email,
			u.name,
			d.id AS device_id,
			d.model,
			d.serial_number,
			d.device_type,
			d.install_status
		FROM relations r
		LEFT JOIN users u ON r.user_id = u.id
		LEFT JOIN devices d ON r.device_id = d.id
		WHERE r.user_id = $1
		ORDER BY r.start_date DESC
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []relations.RelationWithDetails
	for rows.Next() {
		var rel relations.RelationWithDetails
		if err := rows.Scan(
			&rel.ID,
			&rel.StartDate,
			&rel.EndDate,
			&rel.User.ID,
			&rel.User.Email,
			&rel.User.Name,
			&rel.Device.ID,
			&rel.Device.Model,
			&rel.Device.SerialNumber,
			&rel.Device.DeviceType,
			&rel.Device.InstallStatus,
		); err != nil {
			return nil, err
		}
		list = append(list, rel)
	}
	return list, nil
}

// --- CREATE ---
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

// --- END RELATION ---
func End(db *sql.DB, id string) error {
	query := `UPDATE relations SET end_date = $1 WHERE id = $2`
	_, err := db.Exec(query, time.Now(), id)
	return err
}

// --- CHECK ACTIVE ---
func HasActiveRelation(db *sql.DB, deviceID string) (bool, error) {
	query := `
		SELECT 1 FROM relations 
		WHERE device_id = $1 AND end_date IS NULL 
		LIMIT 1
	`
	var exists int
	err := db.QueryRow(query, deviceID).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}