package controller

import (
	"shoplink/app/constant"
	"shoplink/app/domain/dao"
	"shoplink/app/pkg"
	"shoplink/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAllProducts(C *gin.Context)
	GetProductByID(C *gin.Context)
	GetProductByStoreID(C *gin.Context)
	CreateProduct(C *gin.Context)
	UpdateProduct(C *gin.Context)
	DeleteProduct(C *gin.Context)
}

type ProductControllerImpl struct {
	service service.ProductService
}

func ProductControllerInit(service service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		service: service,
	}
}

func (p *ProductControllerImpl) GetAllProducts(c *gin.Context) {
	products, err := p.service.GetAllProducts()
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(200, pkg.BuildResponse(constant.Success, products))
}

func (p *ProductControllerImpl) GetProductByID(c *gin.Context) {
	productid := c.Param("id")
	id, err := strconv.ParseUint(productid, 10, 64)
	if err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid product ID")
	}

	product, err := p.service.GetProductByID(uint(id))
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(200, pkg.BuildResponse(constant.Success, product))
}

func (p *ProductControllerImpl) GetProductByStoreID(c *gin.Context) {
	storeid := c.Param("store_id")
	storeID, err := strconv.ParseUint(storeid, 10, 64)
	if err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid store ID")
	}

	products, err := p.service.GetProductByStoreID(uint(storeID))
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(200, pkg.BuildResponse(constant.Success, products))
}

func (p *ProductControllerImpl) CreateProduct(c *gin.Context) {
	var product dao.Product

	if err := c.ShouldBind(&product); err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid request body")
	}

	created, err := p.service.CreateProduct(c.Request.Context(), product, c.Request.MultipartForm.File["images"], "product-images")
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(200, pkg.BuildResponse(constant.Success, created))
}

func (p *ProductControllerImpl) UpdateProduct(c *gin.Context) {
	var product dao.Product

	if err := c.ShouldBind(&product); err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid request body")
	}

	deletedImageIdsStr := c.PostFormArray("deleted_image_ids")
	var deletedImageIDs []uint
	for _, idStr := range deletedImageIdsStr {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid deleted image ID")
		}
		deletedImageIDs = append(deletedImageIDs, uint(id))
	}

	updated, err := p.service.UpdateProduct(c.Request.Context(), product, c.Request.MultipartForm.File["images"], "product-images", deletedImageIDs)
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(200, pkg.BuildResponse(constant.Success, updated))
}

func (p *ProductControllerImpl) DeleteProduct(c *gin.Context) {
	productid := c.Param("id")
	id, err := strconv.ParseUint(productid, 10, 64)
	if err != nil {
		pkg.PanicException_(constant.InvalidRequest.GetResponseStatus(), "Invalid product ID")
	}

	err = p.service.DeleteProduct(uint(id))
	if err != nil {
		pkg.PanicException_(constant.UnknownError.GetResponseStatus(), err.Error())
	}
	c.JSON(200, pkg.BuildResponse(constant.Success, "Product deleted successfully"))
}
