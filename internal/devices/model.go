package devices

type Device struct {
	ID            string `json:"id"`
	DeviceType    string `json:"device_type"`
	SerialNumber  string `json:"serial_number"`
	Model         string `json:"model"`
	OrderID       string `json:"order_id"`
	InstallStatus string `json:"install_status"`
	CreatedAt     string `json:"created_at,omitempty"`
}