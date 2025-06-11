package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"gorm.io/gorm"
)

type DeathCertificateDocumentRepository interface {
	Create(tx *gorm.DB, DeathDocument *entity.DeathCertificateDocument) error
	Update(tx *gorm.DB, DeathDocument *entity.DeathCertificateDocument) error
	Delete(tx *gorm.DB, DeathDocument *entity.DeathCertificateDocument) error
	FindById(tx *gorm.DB, DeathDocument *entity.DeathCertificateDocument) error
	FindAll(tx *gorm.DB, DeathDocument *[]entity.DeathCertificateDocument) error
}

type DeathCertificateDocumentRepositoryImpl struct {
	Repository[entity.DeathCertificateDocument]
}

func NewDeathCertificateDocumentRepository() DeathCertificateDocumentRepository {
	return &DeathCertificateDocumentRepositoryImpl{}
}

// FindAll implements DeathCertificateDocumentRepository.
func (repository *DeathCertificateDocumentRepositoryImpl) FindAll(tx *gorm.DB, deathDocuments *[]entity.DeathCertificateDocument) error {
	return tx.Find(deathDocuments).Error
}

// FindById implements DeathCertificateDocumentRepository.
func (repository *DeathCertificateDocumentRepositoryImpl) FindById(tx *gorm.DB, deathDocuments *entity.DeathCertificateDocument) error {
	return tx.First(deathDocuments).Error
}
