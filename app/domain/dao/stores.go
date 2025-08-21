package dao

type Store struct {
	BaseModel
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Address     string `gorm:"not null" json:"address"`
	Phone       string `gorm:"not null" json:"phone"`
	UserID      uint   `gorm:"not null" json:"user_id"` // ID of the user who owns the store
	Products 	[]Product `gorm:"foreignKey:StoreID" json:"products"` // products associated with the store       
}
