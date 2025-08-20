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

	r := gin.Default()
	r.Run()
}