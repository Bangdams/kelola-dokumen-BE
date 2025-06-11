package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"gorm.io/gorm"
)

type RwListRepository interface {
	Create(tx *gorm.DB, rwList *entity.RwList) error
	Update(tx *gorm.DB, rwList *entity.RwList) error
	Delete(tx *gorm.DB, rwList *entity.RwList) error
	FindByName(tx *gorm.DB, RwList *entity.RwList) error
}

type RwListRepositoryImpl struct {
	Repository[entity.RwList]
}

func NewRwListRepository() RwListRepository {
	return &RwListRepositoryImpl{}
}

// FindByName implements UserRepository.
func (repository *RwListRepositoryImpl) FindByName(tx *gorm.DB, rwList *entity.RwList) error {
	return tx.First(rwList, "name_rw=?", rwList.NameRw).Error
}
