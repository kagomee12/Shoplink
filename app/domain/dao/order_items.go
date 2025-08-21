package dao

type OrderItem struct {
	BaseModel
	OrderID   uint   `gorm:"not null" json:"order_id"`
	ProductID uint   `gorm:"not null" json:"product_id"`
	Quantity  int    `gorm:"not null" json:"quantity"`
	Price     string `gorm:"not null" json:"price"` 
}