package main

import (
	"shoplink/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.InitLog()
	config.ConnectDB()
	config.Migrate()

	r := gin.Default()
	r.Run()
}