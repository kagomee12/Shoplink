package dao

type Category struct {
	BaseModel
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	ParentID    uint   `gorm:"default:0" json:"parent_id"` // 0 for top-level categories
}