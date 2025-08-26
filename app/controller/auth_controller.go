package controller

import (
	"shoplink/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type AuthControllerImpl struct {
	service service.AuthService
}

func AuthControllerInit (service service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		service: service,
	}
}

func (a *AuthControllerImpl) Login(c *gin.Context) {
	a.service.Login(c)
}

// Register registers a new user to the system.
//
// It takes the user registration data in the request body. If the request body
// is invalid, it returns a 400 error. If there is an error creating the user,
// it returns a 500 error. If the user is created successfully, it returns a 200
// status with the created user data.
func (a *AuthControllerImpl) Register(c *gin.Context) {
	a.service.Register(c)
}

func (a *AuthControllerImpl) RefreshToken(c *gin.Context) {
	a.service.RefreshToken(c)
}