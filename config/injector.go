// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"shoplink/app/config"
	"shoplink/app/controller"
	"shoplink/app/pkg"
	"shoplink/app/repository"
	"shoplink/app/service"

	"github.com/google/wire"
)

var db = wire.NewSet(ConnectDB)

var minioConfig = wire.NewSet(
	config.NewMinioConfig,
)

var userRepo = wire.NewSet(
	repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var minioRepo = wire.NewSet(
	repository.MinioRepositoryInit,
	wire.Bind(new(repository.MinioRepository), new(*repository.MinioRepositoryImpl)),
)

var productRepo = wire.NewSet(
	repository.ProductRepositoryInit,
	wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)),
)

var authService = wire.NewSet(
	service.NewAuthService,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
)

var productService = wire.NewSet(
	service.NewProductService,
	wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)),
)

var authController = wire.NewSet(
	controller.AuthControllerInit,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

var productController = wire.NewSet(
	controller.ProductControllerInit,
	wire.Bind(new(controller.ProductController), new(*controller.ProductControllerImpl)),
)

var jwt = wire.NewSet(
	pkg.NewJWTService,
	wire.Bind(new(pkg.JWTService), new(*pkg.JWTServiceImpl)),
	pkg.NewJWTSecret,
	pkg.NewJWTIssuer,
)

func Init() *Initialization {
	wire.Build(
		InitAll,
		db,
		minioConfig,
		userRepo,
		authService,
		productService,
		authController,
		minioRepo,
		productRepo,
		productController,
		jwt,
	)
	return nil
}
