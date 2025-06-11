package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"gorm.io/gorm"
)

type MarriageStatementDocumentRepository interface {
	Create(tx *gorm.DB, marriageDocument *entity.MarriageStatementDocument) error
	Update(tx *gorm.DB, marriageDocument *entity.MarriageStatementDocument) error
	Delete(tx *gorm.DB, marriageDocument *entity.MarriageStatementDocument) error
	FindById(tx *gorm.DB, marriageDocument *entity.MarriageStatementDocument) error
	FindAll(tx *gorm.DB, marriageDocument *[]entity.MarriageStatementDocument) error
}

type MarriageStatementDocumentRepositoryImpl struct {
	Repository[entity.MarriageStatementDocument]
}

func NewMarriageStatementDocumentRepository() MarriageStatementDocumentRepository {
	return &MarriageStatementDocumentRepositoryImpl{}
}

// FindAll implements MarriageStatementDocumentRepository.
func (repository *MarriageStatementDocumentRepositoryImpl) FindAll(tx *gorm.DB, marriageDocuments *[]entity.MarriageStatementDocument) error {
	return tx.Find(marriageDocuments).Error
}

// FindById implements MarriageStatementDocumentRepository.
func (repository *MarriageStatementDocumentRepositoryImpl) FindById(tx *gorm.DB, marriageDocuments *entity.MarriageStatementDocument) error {
	return tx.First(marriageDocuments).Error
}
