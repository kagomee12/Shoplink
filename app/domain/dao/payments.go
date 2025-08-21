package dao

type Payment struct {
	BaseModel
	OrderID       uint   `gorm:"not null" json:"order_id"`
	PayemntMethod string `gorm:"not null" json:"payment_method"` // e.g., credit card, PayPal
	TransactionID string `gorm:"not null" json:"transaction_id"`
	Status        string `gorm:"not null,default:'pending'" json:"status"` // e.g., waiting, completed, failed
	Amount        string `gorm:"not null" json:"amount"`                   // total amount paid
	PaymentURL    string `gorm:"not null" json:"payment_url"`              // URL for payment gateway
	VaNumber      string `gorm:"not null" json:"va_number"`                // Virtual Account Number for bank transfers
	QrCodeURL     string `gorm:"not null" json:"qr_code_url"`              // URL for QR code payment
}
