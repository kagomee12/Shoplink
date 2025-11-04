package repository

import (
	"shoplink/app/domain/dao"

	"gorm.io/gorm"
)

type StoreRepository interface {
	GetAllStores() ([]dao.Store, error)
	GetStoreByID(id uint) (dao.Store, error)
	CreateStore(store dao.Store) (dao.Store, error)
	UpdateStore(store dao.Store, id uint) (dao.Store, error)
	DeleteStore(id uint) error
}

type StoreRepositoryImpl struct {
	db *gorm.DB
}

func StoreRepositoryInit(db *gorm.DB) *StoreRepositoryImpl {
	return &StoreRepositoryImpl{
		db: db,
	}
}

func (r *StoreRepositoryImpl) GetAllStores() ([]dao.Store, error) {
	var stores []dao.Store
	if err := r.db.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}

func (r *StoreRepositoryImpl) GetStoreByID(id uint) (dao.Store, error) {
	var store dao.Store
	if err := r.db.Preload("Products").First(&store, id).Error; err != nil {
		return dao.Store{}, err
	}
	return store, nil
}

func (r *StoreRepositoryImpl) CreateStore(store dao.Store) (dao.Store, error) {
	if err := r.db.Create(&store).Error; err != nil {
		return dao.Store{}, err
	}
	return store, nil
}

func (r *StoreRepositoryImpl) UpdateStore(store dao.Store, id uint) (dao.Store, error) {
	if err := r.db.Model(&dao.Store{}).Where("ID = ?", id).Updates(store).Error; err != nil {
		return dao.Store{}, err
	}
	return store, nil
}

func (r *StoreRepositoryImpl) DeleteStore(id uint) error {
	if err := r.db.Delete(&dao.Store{}, id).Error; err != nil {
		return err
	}
	return nil
}
