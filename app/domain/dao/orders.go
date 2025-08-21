package dao

type Order struct {
	BaseModel
	UserID            uint        `gorm:"not null" json:"user_id"`
	TotalAmount       string      `gorm:"not null" json:"total_amount"`
	Status            string      `gorm:"not null" json:"status"` // pending, paid, shipped, completed, cancelled
	PaymentID         uint        `gorm:"not null" json:"payment_id"`
	ShippingAddressID string      `gorm:"not null" json:"shipping_address"`
	OrderItems        []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"` // items in the order
	ShippingAddress   Address     `gorm:"foreignKey:ShippingAddressID"`          // address where the order will be shipped
	Payment           Payment     `gorm:"foreignKey:PaymentID"`                  // payment details for the order
}
