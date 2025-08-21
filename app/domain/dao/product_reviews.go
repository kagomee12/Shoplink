package dao

type ProductReview struct {
	BaseModel
	ProductID uint   `gorm:"not null" json:"store_id"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	Rating    int    `gorm:"not null" json:"rating"`
	Comment   string `gorm:"not null" json:"comment"`
}
