package dao

type Product struct {
	BaseModel
	StoreID        uint           `gorm:"not null" json:"store_id"`
	Name           string         `gorm:"not null" json:"name"`
	Description    string         `gorm:"not null" json:"description"`
	Price          string         `gorm:"not null" json:"price"`
	Stock          int            `gorm:"not null" json:"stock"`
	IsActive       bool           `gorm:"default:true" json:"is_active"`
	CategoryID     uint           `gorm:"not null" json:"category_id"`
	ProductImages  []ProductImage `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"product_images"` // images associated with the product
	Category       Category
	ProductReviews []ProductReview `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"product_reviews"` // reviews for the product
}
