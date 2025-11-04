package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/superstan777/stock-backend/internal/devices"
)

// --- GET all devices ---
func GetAllDevices(db *sql.DB) ([]devices.Device, error) {
	rows, err := db.Query(`
		SELECT id, device_type, serial_number, model, order_id, install_status, created_at
		FROM devices
		ORDER BY device_type, serial_number
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []devices.Device
	for rows.Next() {
		var d devices.Device
		if err := rows.Scan(
			&d.ID,
			&d.DeviceType,
			&d.SerialNumber,
			&d.Model,
			&d.OrderID,
			&d.InstallStatus,
			&d.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, d)
	}
	return result, rows.Err()
}


// --- GET devices by type + filters + pagination ---
func GetDevicesByType(db *sql.DB, deviceType string, filters map[string][]string, page, perPage int) ([]devices.Device, int, error) {
	var whereClauses []string
	args := []interface{}{deviceType}
	whereClauses = append(whereClauses, fmt.Sprintf("device_type = $%d", len(args)))

	argIndex := len(args)

	for key, values := range filters {
		if len(values) == 0 {
			continue
		}
		orParts := []string{}
		for _, v := range values {
			argIndex++
			orParts = append(orParts, fmt.Sprintf("%s ILIKE $%d", key, argIndex))
			args = append(args, v+"%")
		}
		whereClauses = append(whereClauses, "("+strings.Join(orParts, " OR ")+")")
	}

	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	offset := (page - 1) * perPage
	query := fmt.Sprintf(`
		SELECT id, device_type, serial_number, model, order_id, install_status, created_at
		FROM devices
		%s
		ORDER BY serial_number
		LIMIT %d OFFSET %d
	`, whereSQL, perPage, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var result []devices.Device
	for rows.Next() {
		var d devices.Device
		if err := rows.Scan(
			&d.ID,
			&d.DeviceType,
			&d.SerialNumber,
			&d.Model,
			&d.OrderID,
			&d.InstallStatus,
			&d.CreatedAt,
		); err != nil {
			return nil, 0, err
		}
		result = append(result, d)
	}

	// count total
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM devices %s", whereSQL)
	row := db.QueryRow(countQuery, args...)
	var total int
	if err := row.Scan(&total); err != nil {
		return nil, 0, err
	}

	return result, total, nil
}


// --- GET one ---
func GetDeviceByID(db *sql.DB, id string) (*devices.Device, error) {
	row := db.QueryRow(`
		SELECT id, device_type, serial_number, model, order_id, install_status, created_at
		FROM devices WHERE id = $1
	`, id)

	var d devices.Device
	if err := row.Scan(
		&d.ID,
		&d.DeviceType,
		&d.SerialNumber,
		&d.Model,
		&d.OrderID,
		&d.InstallStatus,
		&d.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}


// --- INSERT ---
func CreateDevice(db *sql.DB, d *devices.Device) error {
	_, err := db.Exec(`
		INSERT INTO devices (device_type, serial_number, model, order_id, install_status)
		VALUES ($1, $2, $3, $4, $5)
	`, d.DeviceType, d.SerialNumber, d.Model, d.OrderID, d.InstallStatus)
	return err
}


// --- UPDATE ---
func UpdateDevice(db *sql.DB, id string, d *devices.Device) error {
	_, err := db.Exec(`
		UPDATE devices
		SET device_type = $1,
			serial_number = $2,
			model = $3,
			order_id = $4,
			install_status = $5
		WHERE id = $6
	`, d.DeviceType, d.SerialNumber, d.Model, d.OrderID, d.InstallStatus, id)
	return err
}


// --- DELETE ---
func DeleteDevice(db *sql.DB, id string) error {
	_, err := db.Exec(`DELETE FROM devices WHERE id = $1`, id)
	return err
}