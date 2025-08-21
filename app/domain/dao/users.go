package dao

type User struct {
	BaseModel
	Email          string `gorm:"unique;not null" json:"email"`
	Password       string `gorm:"not null" json:"password"`
	Name           string `gorm:"not null" json:"name"`
	Phone          string `gorm:"not null" json:"phone"`
	Role           string `gorm:"not null, default:'buyer'" json:"role"` // buyer, seller, admin
	Store          Store `gorm:"foreignKey:UserID" json:"store"` // store associated with the user, if any
	Addresses      []Address `gorm:"foreignKey:UserID" json:"addresses"` // addresses associated with the user
	Cart           Cart `gorm:"foreignKey:UserID" json:"cart"` // user's shopping cart
	Orders         []Order `gorm:"foreignKey:UserID" json:"orders"` // orders placed by the user
	ProductReviews []ProductReview `gorm:"foreignKey:UserID" json:"product_reviews"` // product reviews written by the user
}
