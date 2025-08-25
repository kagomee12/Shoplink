package config

import (
	"shoplink/app/controller"
	"shoplink/app/pkg"
	"shoplink/app/repository"
	"shoplink/app/service"
)

type Initialization struct {
	userRepo       repository.UserRepository
	AuthService    service.AuthService
	AuthController controller.AuthController
	Jwt            pkg.JWTService
}

func InitAll(
	userRepo repository.UserRepository,
	authService service.AuthService,
	authController controller.AuthController,
	jwt pkg.JWTService,
	) *Initialization {
	return &Initialization{
		userRepo:       userRepo,
		AuthService:    authService,
		AuthController: authController,
		Jwt:            jwt,
	}
}
