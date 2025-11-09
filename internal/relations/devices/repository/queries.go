package repository

import (
	"database/sql"

	"github.com/superstan777/stock-backend/internal/relations"
)

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

	list := []relations.RelationWithDetails{} // zawsze inicjalizowana

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