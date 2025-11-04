package service

import (
	"shoplink/app/domain/dao"
	"shoplink/app/repository"
)

type StoreService interface {
	GetAllStores() ([]dao.Store, error)
	GetStoreByID(id uint) (dao.Store, error)
	CreateStore(store dao.Store) (dao.Store, error)
	UpdateStore(store dao.Store, id uint) (dao.Store, error)
	DeleteStore(id uint) error
}

type StoreServiceImpl struct {
	repo repository.StoreRepository
}

func NewStoreService(repo repository.StoreRepository) *StoreServiceImpl {
	return &StoreServiceImpl{
		repo: repo,
	}
}

func (s *StoreServiceImpl) GetAllStores() ([]dao.Store, error) {
	return s.repo.GetAllStores()
}

func (s *StoreServiceImpl) GetStoreByID(id uint) (dao.Store, error) {
	return s.repo.GetStoreByID(id)
}

func (s *StoreServiceImpl) CreateStore(store dao.Store) (dao.Store, error) {
	return s.repo.CreateStore(store)
}

func (s *StoreServiceImpl) UpdateStore(store dao.Store, id uint) (dao.Store, error) {
	return s.repo.UpdateStore(store, id)
}

func (s *StoreServiceImpl) DeleteStore(id uint) error {
	return s.repo.DeleteStore(id)
}
