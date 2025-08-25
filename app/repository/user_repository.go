package repository

import (
	"shoplink/app/domain/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUsers() ([]dao.User, error)
	FindUserByID(id uint) (dao.User, error)
	CreateUser(user dao.User) (dao.User, error)
	UpdateUser(user dao.User) (dao.User, error)
	DeleteUser(id uint) error
}

type UserRepositoryImpl struct{
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) FindAllUsers() ([]dao.User, error) {
	var users []dao.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindUserByID(id uint) (dao.User, error) {
	user := dao.User{
		BaseModel: dao.BaseModel{ID: id},
	}
	
	if err := u.db.First(&user, id).Error; err != nil {
		return dao.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) CreateUser(user dao.User) (dao.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return dao.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) UpdateUser(user dao.User) (dao.User, error) {
	if err := u.db.Save(&user).Error; err != nil {
		return dao.User{}, err
	}
	return user, nil
}

func (u *UserRepositoryImpl) DeleteUser(id uint) error {
	user := dao.User{
		BaseModel: dao.BaseModel{ID: id},
	}

	if err := u.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}