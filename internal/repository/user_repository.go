package repository

import (
	"fmt"
	"log"

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
	DashboardAdmin(tx *gorm.DB, totalCompleted *int64, totalIncomplete *int64) error
}

type UserRepositoryImpl struct {
	Repository[entity.User]
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

// DashboardAdmin implements UserRepository.
func (repository *UserRepositoryImpl) DashboardAdmin(tx *gorm.DB, totalCompleted *int64, totalIncomplete *int64) error {
	var counts []int64
	var countsCompletes []int64
	tables := []string{
		"mails",
		"death_certificate_documents",
		"marriage_statement_documents",
		"birth_certificate_documents",
	}

	for _, table := range tables {
		var count int64
		err := tx.Table(table).Where("status = ?", "belum").Count(&count).Error
		if err != nil {
			log.Fatal(err)
			return err
		}
		counts = append(counts, count)

		var countsComplete int64
		err = tx.Table(table).Where("status = ?", "selesai").Count(&countsComplete).Error
		if err != nil {
			log.Fatal(err)
			return err
		}
		countsCompletes = append(countsCompletes, countsComplete)
	}

	// Menjumlahkan semua hasil
	for _, count := range counts {
		*totalIncomplete += count
	}

	for _, count := range countsCompletes {
		*totalCompleted += count
	}

	fmt.Println("Total dokumen yang selesai:", totalCompleted)
	fmt.Println("Total dokumen yang belum selesai:", totalIncomplete)

	return nil
}

// FindByUsername implements UserRepository.
func (repository *UserRepositoryImpl) FindByUsername(tx *gorm.DB, user *entity.User) error {
	return tx.First(user, "username=?", user.Username).Error
}

// FindAll implements UserRepository.
func (repository *UserRepositoryImpl) FindAll(tx *gorm.DB, userId uint, users *[]entity.User) error {
	return tx.Preload("RwList").Not("id = ?", userId).Find(users).Error
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
