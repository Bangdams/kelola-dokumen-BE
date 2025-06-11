package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"gorm.io/gorm"
)

type BirthCertificateDocumentRepository interface {
	Create(tx *gorm.DB, birthDocument *entity.BirthCertificateDocument) error
	Update(tx *gorm.DB, birthDocument *entity.BirthCertificateDocument) error
	Delete(tx *gorm.DB, birthDocument *entity.BirthCertificateDocument) error
	FindById(tx *gorm.DB, birthDocument *entity.BirthCertificateDocument) error
	FindAll(tx *gorm.DB, birthDocument *[]entity.BirthCertificateDocument) error
}

type BirthCertificateDocumentRepositoryImpl struct {
	Repository[entity.BirthCertificateDocument]
}

func NewBirthCertificateDocumentRepository() BirthCertificateDocumentRepository {
	return &BirthCertificateDocumentRepositoryImpl{}
}

// FindAll implements BirthCertificateDocumentRepository.
func (repository *BirthCertificateDocumentRepositoryImpl) FindAll(tx *gorm.DB, birthDocuments *[]entity.BirthCertificateDocument) error {
	return tx.Find(birthDocuments).Error
}

// FindById implements BirthCertificateDocumentRepository.
func (repository *BirthCertificateDocumentRepositoryImpl) FindById(tx *gorm.DB, birthDocuments *entity.BirthCertificateDocument) error {
	return tx.First(birthDocuments).Error
}
