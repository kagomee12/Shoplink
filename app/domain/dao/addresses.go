package dao

type Address struct {
	BaseModel
	UserID          uint   `gorm:"not null" json:"user_id"`
	City            string `gorm:"not null" json:"city"`
	State           string `gorm:"not null" json:"state"`
	PostalCode      string `gorm:"not null" json:"postal_code"`
	Country         string `gorm:"not null" json:"country"`
	CompleteAddress string `gorm:"not null" json:"complete_address"` // full address including street, house number, etc.
	IsDefault       bool   `gorm:"default:false" json:"is_default"`  // true if this is the default address for the user
}
