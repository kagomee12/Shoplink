// go:build wireinject
//go:build wireinject
// +build wireinject

package config

import (
	"shoplink/app/controller"
	"shoplink/app/pkg"
	"shoplink/app/repository"
	"shoplink/app/service"

	"github.com/google/wire"
)


var db = wire.NewSet(ConnectDB)

var userRepo = wire.NewSet(
	repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var authService = wire.NewSet(
	service.NewAuthService,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
)

var authController = wire.NewSet(
	controller.AuthControllerInit,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
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
		userRepo,
		authService,
		authController,
		jwt,
	)
	return nil
}