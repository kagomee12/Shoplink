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

type StoreController interface {
	GetAllStores(c *gin.Context)
	GetStoreByID(c *gin.Context)
	CreatedStore(c *gin.Context)
	UpdateStore(c *gin.Context)
	DeleteStore(c *gin.Context)
}

type StoreControllerImpl struct {
	service service.StoreService
}

func StoreControllerInit(service service.StoreService) *StoreControllerImpl {
	return &StoreControllerImpl{
		service: service,
	}
}

func (s *StoreControllerImpl) GetAllStores(c *gin.Context) {
	stores, err := s.service.GetAllStores()

	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, stores))
}

func (s *StoreControllerImpl) GetStoreByID(c *gin.Context) {
	storeId := c.Param("id")

	id, err := strconv.ParseUint(storeId, 10, 64)

	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
		return
	}

	store, err := s.service.GetStoreByID(uint(id))

	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, store))
}

func (s *StoreControllerImpl) CreatedStore(c *gin.Context) {
	var store dao.Store

	if err := c.ShouldBindJSON(&store); err != nil {
		pkg.PanicException(constant.InvalidRequest)
		return
	}

	userID, ok := c.Get("userId")

	if !ok {
		pkg.PanicException(constant.Unauthorized)
		return
	}

	store.UserID = userID.(uint)

	result, err := s.service.CreateStore(store)

	if err != nil {
		pkg.PanicException(constant.UnknownError)
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, result))

}

func (s *StoreControllerImpl) UpdateStore(c *gin.Context) {
	var store dao.Store

	if err := c.ShouldBindJSON(&store); err != nil {
		pkg.PanicException(constant.InvalidRequest)
		return
	}

	storeId := c.Param("storeId")

	idStore, _ := strconv.ParseUint(storeId, 10, 64)

	store.ID = uint(idStore)

	result, err := s.service.UpdateStore(store, uint(idStore))

	if err != nil {
		pkg.PanicException(constant.UnknownError)
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, result))
}

func (s *StoreControllerImpl) DeleteStore(c *gin.Context) {
	storeId := c.Param("storeId")

	idStore, _ := strconv.ParseUint(storeId, 10, 64)

	if err := s.service.DeleteStore(uint(idStore)); err != nil {
		pkg.PanicException(constant.UnknownError)
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "Store deleted successfully"))

}
