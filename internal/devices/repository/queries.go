package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/superstan777/stock-backend/internal/devices"
)

// GetDevices pobiera urządzenia z filtrami i paginacją.
func GetDevices(db *sql.DB, deviceType string, filters map[string]string, page int) ([]devices.Device, int, error) {
	const perPage = 20 // stała liczba wyników na stronę

	// --- PODSTAWOWE ZAPYTANIE ---
	baseQuery := `
		SELECT id, device_type, serial_number, model, order_id, install_status, created_at
		FROM devices
		WHERE 1=1
	`
	args := []interface{}{}
	argIdx := 1

	// --- opcjonalny filter device_type ---
	if deviceType != "" {
		baseQuery += fmt.Sprintf(" AND device_type = $%d", argIdx)
		args = append(args, deviceType)
		argIdx++
	}

	// --- DODATKOWE FILTRY ---
	for key, value := range filters {
		if value == "" {
			continue
		}
		values := strings.Split(value, ",")
		for i := range values {
			values[i] = strings.TrimSpace(values[i])
		}
		if len(values) == 0 {
			continue
		}

		if key == "install_status" {
			placeholders := []string{}
			for _, v := range values {
				placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
				args = append(args, v)
				argIdx++
			}
			baseQuery += fmt.Sprintf(" AND install_status IN (%s)", strings.Join(placeholders, ","))
		} else {
			baseQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, argIdx)
			args = append(args, values[0]+"%")
			argIdx++
		}
	}

	// --- SORTOWANIE ---
	baseQuery += " ORDER BY device_type ASC, serial_number ASC"

	// --- PAGINACJA ---
	offset := (page - 1) * perPage
	baseQuery += fmt.Sprintf(" LIMIT %d OFFSET %d", perPage, offset)

	// --- WYKONANIE ZAPYTANIA ---
	rows, err := db.Query(baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var devicesList []devices.Device
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
		devicesList = append(devicesList, d)
	}

	// --- LICZENIE WSZYSTKICH PASUJĄCYCH URZĄDZEŃ ---
	countQuery := "SELECT COUNT(*) FROM devices WHERE 1=1"
	countArgs := []interface{}{}
	argIdx = 1

	if deviceType != "" {
		countQuery += fmt.Sprintf(" AND device_type = $%d", argIdx)
		countArgs = append(countArgs, deviceType)
		argIdx++
	}

	for key, value := range filters {
		if value == "" {
			continue
		}
		values := strings.Split(value, ",")
		for i := range values {
			values[i] = strings.TrimSpace(values[i])
		}
		if len(values) == 0 {
			continue
		}

		if key == "install_status" {
			placeholders := []string{}
			for _, v := range values {
				placeholders = append(placeholders, fmt.Sprintf("$%d", argIdx))
				countArgs = append(countArgs, v)
				argIdx++
			}
			countQuery += fmt.Sprintf(" AND install_status IN (%s)", strings.Join(placeholders, ","))
		} else {
			countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, argIdx)
			countArgs = append(countArgs, values[0]+"%")
			argIdx++
		}
	}

	var totalCount int
	if err := db.QueryRow(countQuery, countArgs...).Scan(&totalCount); err != nil {
		return nil, 0, err
	}

	return devicesList, totalCount, nil
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