package repository

import (
	"shoplink/app/domain/dao"

	"gorm.io/gorm"
)

type ImageRepository interface {
	FindAllImages() ([]dao.ProductImage, error)
	FindImageByID(id uint) (dao.ProductImage, error)
	CreateImage(image dao.ProductImage) (dao.ProductImage, error)
	UpdateImage(image dao.ProductImage) (dao.ProductImage, error)
	DeleteImage(id uint) error
}

type ImageRepositoryImpl struct {
	db *gorm.DB
}

func ImageRepositoryInit(db *gorm.DB) *ImageRepositoryImpl {
	return &ImageRepositoryImpl{
		db: db,
	}
}

func (r *ImageRepositoryImpl) FindAllImages() ([]dao.ProductImage, error) {
	var Images []dao.ProductImage
	if err := r.db.Find(&Images).Error; err != nil {
		return nil, err
	}
	return Images, nil
}

func (r *ImageRepositoryImpl) FindImageByID(id uint) (dao.ProductImage, error) {
	var image dao.ProductImage
	if err := r.db.First(&image, id).Error; err != nil {
		return dao.ProductImage{}, err
	}
	return image, nil
}

func (r *ImageRepositoryImpl) CreateImage(image dao.ProductImage) (dao.ProductImage, error) {
	if err := r.db.Create(&image).Error; err != nil {
		return dao.ProductImage{}, err
	}
	return image, nil
}

func (r *ImageRepositoryImpl) UpdateImage(image dao.ProductImage) (dao.ProductImage, error) {
	if err := r.db.Save(&image).Error; err != nil {
		return dao.ProductImage{}, err
	}
	return image, nil
}

func (r *ImageRepositoryImpl) DeleteImage(id uint) error {
	if err := r.db.Delete(&dao.ProductImage{}, id).Error; err != nil {
		return err
	}
	return nil
}
