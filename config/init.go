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
	productService service.ProductService,
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
		productService:    productService,
		AuthController:    authController,
		ProductController: productController,
		Jwt:               jwt,
	}
}
