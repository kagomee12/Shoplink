package router

import (
	"shoplink/app/middleware"
	"shoplink/app/pkg"
	"shoplink/config"

	"github.com/gin-gonic/gin"
)

func Init(init *config.Initialization) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(pkg.PanicHandler())

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", init.AuthController.Register)
			auth.POST("/login", init.AuthController.Login)
		}

		token := api.Group("/token")
		{
			token.Use(middleware.JWTMiddleware(init.Jwt))
			api.POST("/refresh-token", init.AuthController.RefreshToken)
		}
		user := api.Group("/users")
		{
			token.Use(middleware.JWTMiddleware(init.Jwt))
			user.GET("/", init.UserController.GetAllUsers)
			user.GET("/:id", init.UserController.GetUserByID)
			user.POST("/", init.UserController.CreateUser)
			user.PUT("/", init.UserController.UpdateUser)
			user.DELETE("/:id", init.UserController.DeleteUser)
		}
		product := api.Group("/products")
		{
			token.Use(middleware.JWTMiddleware(init.Jwt))
			product.GET("/", init.ProductController.GetAllProducts)
			product.GET("/:id", init.ProductController.GetProductByID)
			product.GET("/store/:storeId", init.ProductController.GetProductByStoreID)
			product.POST("/", init.ProductController.CreateProduct)
			product.PUT("/", init.ProductController.UpdateProduct)
			product.DELETE("/:id", init.ProductController.DeleteProduct)
		}

		store := api.Group("/store")
		{
			token.Use(middleware.JWTMiddleware(init.Jwt))
			store.GET("/", init.StoreController.GetAllStores)
			store.GET("/:id", init.StoreController.GetStoreByID)
			store.POST("/", init.StoreController.CreatedStore)
			store.PUT("/", init.StoreController.UpdateStore)
			store.DELETE("/:id", init.StoreController.DeleteStore)
		}

	}

	return r
}
