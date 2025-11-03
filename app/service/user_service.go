package service

import (
	"shoplink/app/domain/dao"
	"shoplink/app/repository"
)

type UserService interface {
	GetAllUsers() ([]dao.User, error)
	GetUserByID(id uint) (dao.User, error)
	CreateUser(user dao.User) (dao.User, error)
	UpdateUser(user dao.User) (dao.User, error)
	DeleteUser(id uint) error
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (s *UserServiceImpl) GetAllUsers() ([]dao.User, error) {
	return s.repo.FindAllUsers()
}

func (s *UserServiceImpl) GetUserByID(id uint) (dao.User, error) {
	return s.repo.FindUserByID(id)
}

func (s *UserServiceImpl) CreateUser(user dao.User) (dao.User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserServiceImpl) UpdateUser(user dao.User) (dao.User, error) {
	return s.repo.UpdateUser(user)
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
