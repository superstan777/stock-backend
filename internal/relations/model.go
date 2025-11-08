package relations

import "time"

// Relation reprezentuje pojedynczy rekord w tabeli relations
type Relation struct {
	ID        string     `json:"id"`
	DeviceID  string     `json:"device_id"`
	UserID    string     `json:"user_id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

// RelationInsert używany przy dodawaniu nowej relacji (INSERT)
type RelationInsert struct {
	DeviceID  string     `json:"device_id"`
	UserID    string     `json:"user_id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

// RelationWithDetails może być użyty przy JOIN-ach — np. z userem lub device
type RelationWithDetails struct {
	ID        string     `json:"id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`

	Device DeviceInfo `json:"device"`
	User   UserInfo   `json:"user"`
}

// DeviceInfo uproszczona reprezentacja urządzenia — np. do list relacji
type DeviceInfo struct {
	ID           string  `json:"id"`
	Model        string  `json:"model"`
	SerialNumber string  `json:"serial_number"`
	DeviceType   string  `json:"device_type"`
	InstallStatus string `json:"install_status"`
}

// UserInfo uproszczona reprezentacja użytkownika
type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}