package dao

type Cart struct {
	BaseModel
	UserID    uint   `gorm:"not null" json:"user_id"`
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cart_items"` // items in the cart
}