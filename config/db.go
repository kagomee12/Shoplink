package config

import (
	"fmt"
	"log"
	"os"
	"shoplink/app/domain/dao"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(
		&dao.User{},
		&dao.Order{},
		&dao.OrderItem{},
		&dao.ProductReview{},
		&dao.Cart{},
		&dao.CartItem{},
		&dao.Address{},
		&dao.Payment{},
		&dao.Product{},
		&dao.ProductImage{},
		&dao.Store{},
		&dao.ProductImage{},
		&dao.Category{},
	)
	if err != nil {
		log.Fatal("Error migrating database. Error: ", err)
	}
}