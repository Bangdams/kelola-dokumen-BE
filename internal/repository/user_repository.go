package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(tx *gorm.DB, user *entity.User) error
	Update(tx *gorm.DB, user *entity.User) error
	Delete(tx *gorm.DB, user *entity.User) error
	Login(tx *gorm.DB, user *entity.User, keyword string) error
	FindAll(tx *gorm.DB, userId uint, users *[]entity.User) error
	FindById(tx *gorm.DB, user *entity.User) error
	FindByUsername(tx *gorm.DB, user *entity.User) error
}

type UserRepositoryImpl struct {
	Repository[entity.User]
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// FindByUsername implements UserRepository.
func (repository *UserRepositoryImpl) FindByUsername(tx *gorm.DB, user *entity.User) error {
	return tx.First(user, "username=?", user.Username).Error
}

// FindAll implements UserRepository.
func (repository *UserRepositoryImpl) FindAll(tx *gorm.DB, userId uint, users *[]entity.User) error {
	return tx.Not("id = ?", userId).Find(users).Error
}

// FindById implements UserRepository.
func (repository *UserRepositoryImpl) FindById(tx *gorm.DB, user *entity.User) error {
	return tx.First(user).Error
}

// Login implements UserRepository.
func (*UserRepositoryImpl) Login(tx *gorm.DB, user *entity.User, keyword string) error {
	return tx.Where("username = ?", keyword).First(user).Error
}

// Search implements UserRepository.
func (*UserRepositoryImpl) Search(tx *gorm.DB, users *[]entity.User, name string) error {
	return tx.Where("name LIKE ?", "%"+name+"%").Find(&users).Error
}
