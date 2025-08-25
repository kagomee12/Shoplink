package router

import (
	"net/http"
	"shoplink/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		_, err := init.Jwt.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", init.AuthController.Register)
			auth.POST("/login", init.AuthController.Login)
		}
	}

	return r
}