package service

import (
	"context"
	"mime/multipart"
	"path"
	"shoplink/app/domain/dao"
	"shoplink/app/repository"
)

type ProductService interface {
	GetAllProducts() ([]dao.Product, error)
	GetProductByID(id uint) (dao.Product, error)
	GetProductByStoreID(storeID uint) ([]dao.Product, error)
	CreateProduct(ctx context.Context, product dao.Product, files []*multipart.FileHeader, bucketName string) (dao.Product, error)
	UpdateProduct(ctx context.Context, product dao.Product, files []*multipart.FileHeader, bucketName string, deletedImageIDs []uint) (dao.Product, error)
	DeleteProduct(id uint) error
}

type ProductServiceImpl struct {
	repo         repository.ProductRepository
	minio        repository.MinioRepository
	productImage repository.ImageRepository
}

func NewProductService(repo repository.ProductRepository, minio repository.MinioRepository, productImage repository.ImageRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		repo:         repo,
		minio:        minio,
		productImage: productImage,
	}
}

func (s *ProductServiceImpl) GetAllProducts() ([]dao.Product, error) {
	return s.repo.FindAllProducts()
}

func (s *ProductServiceImpl) GetProductByID(id uint) (dao.Product, error) {
	return s.repo.FindProductByID(id)
}

func (s *ProductServiceImpl) GetProductByStoreID(storeID uint) ([]dao.Product, error) {
	return s.repo.FindProductByStoreID(storeID)
}

func (s *ProductServiceImpl) CreateProduct(ctx context.Context, product dao.Product, files []*multipart.FileHeader, bucketName string) (dao.Product, error) {

	for _, fileHeader := range files {
		file, err := fileHeader.Open()

		if err != nil {
			return dao.Product{}, err
		}

		url, err := s.minio.UploadFile(
			ctx,
			bucketName,
			fileHeader.Filename,
			file,
			fileHeader.Size,
			fileHeader.Header.Get("Content-Type"),
		)
		defer file.Close()

		if err != nil {
			return dao.Product{}, err
		}
		productImage := dao.ProductImage{
			ImageURL:  url,
			ProductID: product.ID,
		}
		product.ProductImages = append(product.ProductImages, productImage)
	}

	return s.repo.CreateProduct(product)
}

func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, product dao.Product, files []*multipart.FileHeader, bucketName string, deletedImageIDs []uint) (dao.Product, error) {
	if len(files) > 0 {
		for _, fileHeader := range files {
			file, err := fileHeader.Open()

			if err != nil {
				return dao.Product{}, err
			}

			url, err := s.minio.UploadFile(
				ctx,
				bucketName,
				fileHeader.Filename,
				file,
				fileHeader.Size,
				fileHeader.Header.Get("Content-Type"),
			)
			defer file.Close()

			if err != nil {
				return dao.Product{}, err
			}
			productImage := dao.ProductImage{
				ImageURL:  url,
				ProductID: product.ID,
			}
			product.ProductImages = append(product.ProductImages, productImage)
		}
	}

	if len(deletedImageIDs) > 0 {
		for _, imageID := range deletedImageIDs {
			var image dao.ProductImage
			image, err := s.productImage.FindImageByID(imageID)
			if err != nil {
				return dao.Product{}, err
			}
			if image.ID != 0 {
				objectName := path.Base(image.ImageURL)
				err := s.minio.DeleteFile(ctx, bucketName, objectName)
				if err != nil {
					return dao.Product{}, err
				}
			}
			err = s.productImage.DeleteImage(image.ID)
			if err != nil {
				return dao.Product{}, err
			}
		}
	}

	return s.repo.UpdateProduct(product)
}

func (s *ProductServiceImpl) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}
