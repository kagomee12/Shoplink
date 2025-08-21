package dao

type CartItem struct {
	BaseModel
	CartID    uint    `gorm:"not null" json:"cart_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     string  `gorm:"not null" json:"price"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"` // Product details associated with the cart item
}
