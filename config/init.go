package config

import (
	"shoplink/app/config"
	"shoplink/app/controller"
	"shoplink/app/pkg"
	"shoplink/app/repository"
	"shoplink/app/service"
)

type Initialization struct {
	minioConfig       config.MinioConfig
	userRepo          repository.UserRepository
	MinioRepo         repository.MinioRepository
	ImageRepo         repository.ImageRepository
	ProductRepo       repository.ProductRepository
	AuthService       service.AuthService
	UserService       service.UserService
	UserController    controller.UserController
	productService    service.ProductService
	AuthController    controller.AuthController
	ProductController controller.ProductController
	Jwt               pkg.JWTService
}

func InitAll(
	minioConfig *config.MinioConfig,
	minioRepo repository.MinioRepository,
	userRepo repository.UserRepository,
	imageRepo repository.ImageRepository,
	ProductRepo repository.ProductRepository,
	authService service.AuthService,
	userService service.UserService,
	productService service.ProductService,
	userController controller.UserController,
	authController controller.AuthController,
	productController controller.ProductController,
	jwt pkg.JWTService,
) *Initialization {
	return &Initialization{
		minioConfig:       *minioConfig,
		MinioRepo:         minioRepo,
		userRepo:          userRepo,
		ImageRepo:         imageRepo,
		ProductRepo:       ProductRepo,
		AuthService:       authService,
		UserService:       userService,
		productService:    productService,
		UserController:    userController,
		AuthController:    authController,
		ProductController: productController,
		Jwt:               jwt,
	}
}
