package main

import (
	"shoplink/app/router"
	"shoplink/config"

	"github.com/joho/godotenv"
)

func init(){
	godotenv.Load()
	config.InitLog()
}

func main() {
	
	r := config.Init()

	app := router.Init(r)
	app.Run(":" + "8080")
}