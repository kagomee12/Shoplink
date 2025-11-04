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
	StoreRepo         repository.StoreRepository
	AuthService       service.AuthService
	UserService       service.UserService
	StoreService      service.StoreService
	ProductService    service.ProductService
	UserController    controller.UserController
	StoreController   controller.StoreController
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
	storeRepo repository.StoreRepository,
	ProductRepo repository.ProductRepository,
	authService service.AuthService,
	userService service.UserService,
	storeService service.StoreService,
	productService service.ProductService,
	userController controller.UserController,
	authController controller.AuthController,
	storeController controller.StoreController,
	productController controller.ProductController,
	jwt pkg.JWTService,
) *Initialization {
	return &Initialization{
		minioConfig:       *minioConfig,
		MinioRepo:         minioRepo,
		userRepo:          userRepo,
		ImageRepo:         imageRepo,
		StoreRepo:         storeRepo,
		ProductRepo:       ProductRepo,
		AuthService:       authService,
		UserService:       userService,
		StoreService:      storeService,
		productService:    productService,
		UserController:    userController,
		AuthController:    authController,
		StoreController:   storeController,
		ProductController: productController,
		Jwt:               jwt,
	}
}
