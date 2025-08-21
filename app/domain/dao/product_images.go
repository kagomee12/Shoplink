package dao

type ProductImage struct {
	BaseModel
	ImageURL string `gorm:"not null" json:"image_url"`
	ProductID uint   `gorm:"not null" json:"product_id"`
}