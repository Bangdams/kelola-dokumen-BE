package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model/converter"
	"gorm.io/gorm"
)

type MailRepository interface {
	Create(tx *gorm.DB, mail *entity.Mail) error
	Update(tx *gorm.DB, mail *entity.Mail) error
	Delete(tx *gorm.DB, mail *entity.Mail) error
	FindById(tx *gorm.DB, mail *entity.Mail) error
	FindAll(tx *gorm.DB, mails *[]entity.Mail) error
	FindAllIncompliteMail(tx *gorm.DB, response *model.AllMailResponse) error
	FindAllCompliteMail(tx *gorm.DB, response *model.AllMailResponse) error
	FindAllMailForUser(tx *gorm.DB, response *[]model.AllMailItemForUserResponse, userId uint) error
}

type MailRepositoryImpl struct {
	Repository[entity.Mail]
}

func NewMailRepository() MailRepository {
	return &MailRepositoryImpl{}
}

// FindAllMailForUser implements MailRepository.
func (repository *MailRepositoryImpl) FindAllMailForUser(tx *gorm.DB, response *[]model.AllMailItemForUserResponse, userId uint) error {
	// Mengambil semua data dari tabel mails dengan status "selesai"
	var mails []entity.Mail
	err := tx.Table("mails").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("mails.user_id = ?", userId).Order("mails.created_at DESC").Find(&mails).Error
	if err != nil {
		return err
	}

	converter.MailForUserResponses(response, &mails)

	// Mengambil semua data dari tabel birth_certificate_documents dengan status "selesai"
	var birthDocuments []entity.BirthCertificateDocument
	err = tx.Table("birth_certificate_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("birth_certificate_documents.user_id = ?", userId).Order("birth_certificate_documents.created_at DESC").Find(&birthDocuments).Error
	if err != nil {
		return err
	}

	converter.BirthDocumentForUserResponses(response, &birthDocuments)

	// Mengambil semua data dari tabel death_certificate_documents dengan status "selesai"
	var deathDocuments []entity.DeathCertificateDocument
	err = tx.Table("death_certificate_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("death_certificate_documents.user_id = ?", userId).Order("death_certificate_documents.created_at DESC").Find(&deathDocuments).Error
	if err != nil {
		return err
	}

	converter.DeathDocumentForUserResponses(response, &deathDocuments)

	// Mengambil semua data dari tabel marriage_statement_documents dengan status "selesai"
	var marriageDocuments []entity.MarriageStatementDocument
	err = tx.Table("marriage_statement_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("marriage_statement_documents.user_id = ?", userId).Order("marriage_statement_documents.created_at DESC").Find(&marriageDocuments).Error
	if err != nil {
		return err
	}

	converter.MarriageDocumentForUserResponses(response, &marriageDocuments)

	return nil
}

// FindAllCompliteMail implements MailRepository.
func (repository *MailRepositoryImpl) FindAllCompliteMail(tx *gorm.DB, response *model.AllMailResponse) error {
	// Mengambil semua data dari tabel mails dengan status "selesai"
	var mails []entity.Mail
	err := tx.Table("mails").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("status = ?", "selesai").Find(&mails).Error
	if err != nil {
		return err
	}

	var mailResponses []model.MailWithUserResponse
	converter.MailCompliteResponses(&mailResponses, &mails)
	response.MailResponses = mailResponses

	// Mengambil semua data dari tabel birth_certificate_documents dengan status "selesai"
	var birthDocuments []entity.BirthCertificateDocument
	err = tx.Table("birth_certificate_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("status = ?", "selesai").Find(&birthDocuments).Error
	if err != nil {
		return err
	}

	var birthDocumentResponses []model.BirthCertificateDocumentWithUserResponse
	converter.BirthDocumentCompliteResponses(&birthDocumentResponses, &birthDocuments)

	response.BirthDocumentResponses = birthDocumentResponses

	// Mengambil semua data dari tabel death_certificate_documents dengan status "selesai"
	var deathDocuments []entity.DeathCertificateDocument
	err = tx.Table("death_certificate_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("status = ?", "selesai").Find(&deathDocuments).Error
	if err != nil {
		return err
	}

	var deathDocumentResponses []model.DeathCertificateDocumentWithUserResponse
	converter.DeathDocumentCompliteResponses(&deathDocumentResponses, &deathDocuments)
	response.DeathDocumentResponses = deathDocumentResponses

	// Mengambil semua data dari tabel marriage_statement_documents dengan status "selesai"
	var marriageDocuments []entity.MarriageStatementDocument
	err = tx.Table("marriage_statement_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("status = ?", "selesai").Find(&marriageDocuments).Error
	if err != nil {
		return err
	}

	var marriageDocumentsResponses []model.MarriageStatementDocumentWithUserResponse
	converter.MarriageDocumentsCompliteResponses(&marriageDocumentsResponses, &marriageDocuments)
	response.MarriageDocumentResponses = marriageDocumentsResponses

	return nil
}

// FindAllIncompliteMail implements MailRepository.
func (repository *MailRepositoryImpl) FindAllIncompliteMail(tx *gorm.DB, response *model.AllMailResponse) error {
	// Mengambil semua data dari tabel mails dengan status "ditunggu"
	var mails []entity.Mail
	err := tx.Table("mails").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("status = ?", "ditunggu").Find(&mails).Error
	if err != nil {
		return err
	}

	var mailResponses []model.MailWithUserResponse
	converter.MailIncompliteResponses(&mailResponses, &mails)
	response.MailResponses = mailResponses

	// Mengambil semua data dari tabel birth_certificate_documents dengan status "ditunggu"
	var birthDocuments []entity.BirthCertificateDocument
	err = tx.Table("birth_certificate_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("status = ?", "ditunggu").Find(&birthDocuments).Error
	if err != nil {
		return err
	}

	var birthDocumentResponses []model.BirthCertificateDocumentWithUserResponse
	converter.BirthDocumentIncompliteResponses(&birthDocumentResponses, &birthDocuments)
	response.BirthDocumentResponses = birthDocumentResponses

	// Mengambil semua data dari tabel death_certificate_documents dengan status "ditunggu"
	var deathDocuments []entity.DeathCertificateDocument
	err = tx.Table("death_certificate_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("status = ?", "ditunggu").Find(&deathDocuments).Error
	if err != nil {
		return err
	}

	var deathDocumentResponses []model.DeathCertificateDocumentWithUserResponse
	converter.DeathDocumentIncompliteResponses(&deathDocumentResponses, &deathDocuments)
	response.DeathDocumentResponses = deathDocumentResponses

	// Mengambil semua data dari tabel marriage_statement_documents dengan status "ditunggu"
	var marriageDocuments []entity.MarriageStatementDocument
	err = tx.Table("marriage_statement_documents").Joins("User.RwList").Joins("StatementTypeItem.StatementType").Where("status = ?", "ditunggu").Find(&marriageDocuments).Error
	if err != nil {
		return err
	}

	var marriageDocumentsResponses []model.MarriageStatementDocumentWithUserResponse
	converter.MarriageDocumentsIncompliteResponses(&marriageDocumentsResponses, &marriageDocuments)
	response.MarriageDocumentResponses = marriageDocumentsResponses

	return nil
}

// FindAll implements MailRepository.
func (repository *MailRepositoryImpl) FindAll(tx *gorm.DB, mails *[]entity.Mail) error {
	return tx.Find(mails).Error
}

// FindById implements UserRepository.
func (repository *MailRepositoryImpl) FindById(tx *gorm.DB, mail *entity.Mail) error {
	return tx.First(mail).Error
}
