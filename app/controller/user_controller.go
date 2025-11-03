package controller

import (
	"net/http"
	"shoplink/app/constant"
	"shoplink/app/domain/dao"
	"shoplink/app/pkg"
	"shoplink/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetMe(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserControllerImpl struct {
	service service.UserService
}

func UserControllerInit(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		service: service,
	}
}

func (u *UserControllerImpl) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, users))
}

func (u *UserControllerImpl) GetUserByID(c *gin.Context) {
	userid := c.Param("id")
	id, err := strconv.ParseUint(userid, 10, 64)
	if err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid user ID")
	}

	user, err := u.service.GetUserByID(uint(id))
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, user))
}

func (u *UserControllerImpl) GetMe(c *gin.Context) {
	userID, _ := c.Get("userID")

	user, err := u.service.GetUserByID(userID.(uint))
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, user))
}

func (u *UserControllerImpl) CreateUser(c *gin.Context) {
	var user dao.User

	if err := c.ShouldBindJSON(&user); err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid request body")
	}

	createdUser, err := u.service.CreateUser(user)
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, createdUser))
}

func (u *UserControllerImpl) UpdateUser(c *gin.Context) {
	var user dao.User

	if err := c.ShouldBindJSON(&user); err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid request body")
	}

	updatedUser, err := u.service.UpdateUser(user)
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, updatedUser))
}

func (u *UserControllerImpl) DeleteUser(c *gin.Context) {
	userid := c.Param("id")
	id, err := strconv.ParseUint(userid, 10, 64)
	if err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid user ID")
	}

	if err := u.service.DeleteUser(uint(id)); err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "User deleted successfully"))
}
