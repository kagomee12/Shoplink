package service

import (
	"net/http"
	"shoplink/app/constant"
	"shoplink/app/domain/dao"
	"shoplink/app/pkg"
	"shoplink/app/repository"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type AuthServiceImpl struct {
	repo repository.UserRepository
	jwt pkg.JWTService
}

func NewAuthService(repo repository.UserRepository, jwt pkg.JWTService) *AuthServiceImpl {
	return &AuthServiceImpl{
		repo: repo,
		jwt: jwt,
	}
}

// Register registers a new user to the system.
//
// It takes the user registration data in the request body. If the request body
// is invalid, it returns a 400 error. If there is an error creating the user,
// it returns a 500 error. If the user is created successfully, it returns a 200
// status with the created user data.
func (s *AuthServiceImpl) Register(c *gin.Context) {
	// defer pkg.PanicHandler(c)

	log.Info("Registering new user")
	var user dao.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		log.Error("Invalid request body for user registration: ", err)
		pkg.PanicException(constant.InvalidRequest)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hash)

	user, err := s.repo.CreateUser(user)

	if err != nil {
		log.Error("Error creating user: ", err)
		pkg.PanicException(constant.UnknownError)
		return
	}

	log.Info("User registered successfully: ", user.ID)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, user))
}

// Login logs in a user to their account
//
// It takes an ID and password in the request body. If the ID does not exist,
// it returns a 404 error. If the password is incorrect, it returns a 401 error.
// If the login is successful, it logs the user in and returns a 200 status.
func (s *AuthServiceImpl) Login(c *gin.Context) {
	// defer pkg.PanicHandler(c)

	log.Info("User login attempt")
	var user dao.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		log.Error("Invalid request body for user login: ", err)
		pkg.PanicException(constant.InvalidRequest)
		return
	}

	existingUser, err := s.repo.FindUserByID(user.ID); if err != nil {
		log.Error("User not found: ", err)
		pkg.PanicException(constant.DataNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		log.Error("Incorrect password for user ID: ", user.ID)
		pkg.PanicException(constant.Unauthorized)
		return
	}

	accessToken, refreshToken, err := s.jwt.GenerateToken(existingUser.ID, existingUser.Name)

	if err != nil {
		log.Error("Error generating token for user ID: ", existingUser.ID, " Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	log.Info("User logged in successfully: ", existingUser.ID)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	}))
}

func (s *AuthServiceImpl) RefreshToken(c *gin.Context) {
	// defer pkg.PanicHandler(c)

	log.Info("Refreshing token")

	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindBodyWithJSON(&requestBody); err != nil {
		log.Error("Invalid request body for token refresh: ", err)
		pkg.PanicException(constant.InvalidRequest)
		return
	}

	claims, err := s.jwt.ValidateToken(requestBody.RefreshToken)

	if err != nil {
		log.Error("Invalid refresh token: ", err)
		pkg.PanicException(constant.Unauthorized)
		return
	}

	accessToken, refreshToken, err := s.jwt.GenerateToken(claims.UserID, claims.Username)

	if err != nil {
		log.Error("Error generating token for user ID: ", claims.UserID, " Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	log.Info("Token refreshed successfully for user ID: ", claims.UserID)
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, gin.H{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	}))
}
