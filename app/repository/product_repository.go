package repository

import (
	"shoplink/app/domain/dao"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAllProducts() ([]dao.Product, error)
	FindProductByID(id uint) (dao.Product, error)
	FindProductByStoreID(storeID uint) ([]dao.Product, error)
	CreateProduct(product dao.Product) (dao.Product, error)
	UpdateProduct(product dao.Product) (dao.Product, error)
	DeleteProduct(id uint) error
}

type ProductRepositoryImpl struct {
	db    *gorm.DB
	minio MinioRepository
}

func ProductRepositoryInit(db *gorm.DB, minio MinioRepository) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db:    db,
		minio: minio,
	}
}

func (r *ProductRepositoryImpl) FindAllProducts() ([]dao.Product, error) {
	var products []dao.Product
	if err := r.db.Preload("ProductImages").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryImpl) FindProductByID(id uint) (dao.Product, error) {
	var product dao.Product
	if err := r.db.Preload("ProductImages").First(&product, id).Error; err != nil {
		return dao.Product{}, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) FindProductByStoreID(storeID uint) ([]dao.Product, error) {
	var products []dao.Product
	if err := r.db.Preload("ProductImages").Where("store_id = ?", storeID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepositoryImpl) CreateProduct(product dao.Product) (dao.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return dao.Product{}, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(product dao.Product) (dao.Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return dao.Product{}, err
	}
	return product, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id uint) error {
	if err := r.db.Delete(&dao.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
