package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"gorm.io/gorm"
)

type MailRepository interface {
	Create(tx *gorm.DB, mail *entity.Mail) error
	Update(tx *gorm.DB, mail *entity.Mail) error
	Delete(tx *gorm.DB, mail *entity.Mail) error
	FindById(tx *gorm.DB, mail *entity.Mail) error
	FindAll(tx *gorm.DB, mails *[]entity.Mail) error
}

type MailRepositoryImpl struct {
	Repository[entity.Mail]
}

func NewMailRepository() MailRepository {
	return &MailRepositoryImpl{}
}

// FindAll implements MailRepository.
func (repository *MailRepositoryImpl) FindAll(tx *gorm.DB, mails *[]entity.Mail) error {
	return tx.Find(mails).Error
}

// FindById implements UserRepository.
func (repository *MailRepositoryImpl) FindById(tx *gorm.DB, mail *entity.Mail) error {
	return tx.First(mail).Error
}
