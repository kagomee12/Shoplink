package dao

type ProductImage struct {
	BaseModel
	ImageURL  string `gorm:"not null" json:"image_url" form:"image_url" binding:"required"`
	ProductID uint   `gorm:"not null" json:"product_id" form:"product_id" binding:"required"`
}
