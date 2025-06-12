package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
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
}

type MailRepositoryImpl struct {
	Repository[entity.Mail]
}

func NewMailRepository() MailRepository {
	return &MailRepositoryImpl{}
}

// FindAllCompliteMail implements MailRepository.
func (repository *MailRepositoryImpl) FindAllCompliteMail(tx *gorm.DB, response *model.AllMailResponse) error {
	// Mengambil semua data dari tabel mails dengan status "selesai"
	var mails []entity.Mail
	err := tx.Table("mails").Joins("User").Joins("StatementTypeItem.StatementType").Where("status = ?", "selesai").Find(&mails).Error
	if err != nil {
		return err
	}

	var mailResponses []model.MailWithUserResponse
	for _, mail := range mails {
		mailResponses = append(mailResponses, model.MailWithUserResponse{
			ID:                mail.ID,
			StatementTypeItem: mail.StatementTypeItem.Name,
			StatementType:     mail.StatementTypeItem.StatementType.Name,
			Username:          mail.User.Username,
			Name:              mail.User.Name,
			KtpFilePath:       mail.KtpFilePath,
			KkFilePath:        mail.KkFilePath,
			RtRwFilePath:      mail.RtRwFilePath,
			SuratBalasan:      mail.SuratBalasanFilePath,
		})
	}

	response.MailResponses = mailResponses

	// Mengambil semua data dari tabel birth_certificate_documents dengan status "selesai"
	var birthDocuments []entity.BirthCertificateDocument
	err = tx.Table("birth_certificate_documents").Joins("User").Joins("StatementTypeItem.StatementType").Where("status = ?", "selesai").Find(&birthDocuments).Error
	if err != nil {
		return err
	}

	var birthDocumentResponses []model.BirthCertificateDocumentWithUserResponse
	for _, birthDocument := range birthDocuments {
		birthDocumentResponses = append(birthDocumentResponses, model.BirthCertificateDocumentWithUserResponse{
			ID:                     birthDocument.ID,
			StatementTypeItem:      birthDocument.StatementTypeItem.Name,
			StatementType:          birthDocument.StatementTypeItem.StatementType.Name,
			Username:               birthDocument.User.Username,
			Name:                   birthDocument.User.Name,
			RtRwFilePath:           birthDocument.RtRwFilePath,
			FormulirFilePath:       birthDocument.FormulirFilePath,
			SuratKelahiranFilePath: birthDocument.SuratKelahiranFilePath,
			BukuNikahFilePath:      birthDocument.BukuNikahFilePath,
			KkFilePath:             birthDocument.KkFilePath,
			KtpPelaporFilePath:     birthDocument.KtpPelaporFilePath,
			KtpSaksi1FilePath:      birthDocument.KtpSaksi1FilePath,
			KtpSaksi2FilePath:      birthDocument.KtpSaksi2FilePath,
			KtpOrangTuaFilePath:    birthDocument.KtpOrangTuaFilePath,
			SuratBalasan:           birthDocument.SuratBalasanFilePath,
		})
	}

	response.BirthDocumentResponses = birthDocumentResponses

	// Mengambil semua data dari tabel death_certificate_documents dengan status "selesai"
	var deathDocuments []entity.DeathCertificateDocument
	err = tx.Table("death_certificate_documents").Joins("User").Joins("StatementTypeItem.StatementType").Where("status = ?", "selesai").Find(&deathDocuments).Error
	if err != nil {
		return err
	}

	var deathDocumentResponses []model.DeathCertificateDocumentWithUserResponse
	for _, deathDocument := range deathDocuments {
		deathDocumentResponses = append(deathDocumentResponses, model.DeathCertificateDocumentWithUserResponse{
			ID:                    deathDocument.ID,
			StatementTypeItem:     deathDocument.StatementTypeItem.Name,
			StatementType:         deathDocument.StatementTypeItem.StatementType.Name,
			Username:              deathDocument.User.Username,
			Name:                  deathDocument.User.Name,
			RtRwFilePath:          deathDocument.RtRwFilePath,
			FormulirFilePath:      deathDocument.FormulirFilePath,
			SuratKematianFilePath: deathDocument.SuratKematianFilePath,
			KtpFilePath:           deathDocument.KtpFilePath,
			KkFilePath:            deathDocument.KkFilePath,
			KtpPelaporFilePath:    deathDocument.KtpPelaporFilePath,
			KtpSaksi1FilePath:     deathDocument.KtpSaksi1FilePath,
			KtpSaksi2FilePath:     deathDocument.KtpSaksi2FilePath,
			BukuNikahFilePath:     deathDocument.BukuNikahFilePath,
			SuratBalasan:          deathDocument.SuratBalasanFilePath,
		})
	}

	response.DeathDocumentResponses = deathDocumentResponses

	// Mengambil semua data dari tabel marriage_statement_documents dengan status "selesai"
	var marriageDocuments []entity.MarriageStatementDocument
	err = tx.Table("marriage_statement_documents").Joins("User").Joins("StatementTypeItem.StatementType").Where("status = ?", "selesai").Find(&marriageDocuments).Error
	if err != nil {
		return err
	}

	var marriageDocumentsResponses []model.MarriageStatementDocumentWithUserResponse
	for _, marriageDocument := range marriageDocuments {
		marriageDocumentsResponses = append(marriageDocumentsResponses, model.MarriageStatementDocumentWithUserResponse{
			ID:                marriageDocument.ID,
			StatementTypeItem: marriageDocument.StatementTypeItem.Name,
			StatementType:     marriageDocument.StatementTypeItem.StatementType.Name,
			Username:          marriageDocument.User.Username,
			Name:              marriageDocument.User.Name,
			KkFilePath:        marriageDocument.KkFilePath,
			KtpFilePath:       marriageDocument.KtpFilePath,
			IjazahFilePath:    marriageDocument.IjazahFilePath,
			AktaFilePath:      marriageDocument.AktaFilePath,
			KtpSaksiFilePath:  marriageDocument.KtpSaksiFilePath,
			SuratBalasan:      marriageDocument.SuratBalasanFilePath,
		})
	}

	response.MarriageDocumentResponses = marriageDocumentsResponses

	return nil
}

// FindAllIncompliteMail implements MailRepository.
func (repository *MailRepositoryImpl) FindAllIncompliteMail(tx *gorm.DB, response *model.AllMailResponse) error {
	// Mengambil semua data dari tabel mails dengan status "belum"
	var mails []entity.Mail
	err := tx.Table("mails").Joins("User").Joins("StatementTypeItem.StatementType").Where("status = ?", "belum").Find(&mails).Error
	if err != nil {
		return err
	}

	var mailResponses []model.MailWithUserResponse
	for _, mail := range mails {
		mailResponses = append(mailResponses, model.MailWithUserResponse{
			ID:                mail.ID,
			StatementTypeItem: mail.StatementTypeItem.Name,
			StatementType:     mail.StatementTypeItem.StatementType.Name,
			Username:          mail.User.Username,
			Name:              mail.User.Name,
			KtpFilePath:       mail.KtpFilePath,
			KkFilePath:        mail.KkFilePath,
			RtRwFilePath:      mail.RtRwFilePath,
		})
	}

	response.MailResponses = mailResponses

	// Mengambil semua data dari tabel birth_certificate_documents dengan status "belum"
	var birthDocuments []entity.BirthCertificateDocument
	err = tx.Table("birth_certificate_documents").Joins("User").Joins("StatementTypeItem.StatementType").Where("status = ?", "belum").Find(&birthDocuments).Error
	if err != nil {
		return err
	}

	var birthDocumentResponses []model.BirthCertificateDocumentWithUserResponse
	for _, birthDocument := range birthDocuments {
		birthDocumentResponses = append(birthDocumentResponses, model.BirthCertificateDocumentWithUserResponse{
			ID:                     birthDocument.ID,
			StatementTypeItem:      birthDocument.StatementTypeItem.Name,
			StatementType:          birthDocument.StatementTypeItem.StatementType.Name,
			Username:               birthDocument.User.Username,
			Name:                   birthDocument.User.Name,
			RtRwFilePath:           birthDocument.RtRwFilePath,
			FormulirFilePath:       birthDocument.FormulirFilePath,
			SuratKelahiranFilePath: birthDocument.SuratKelahiranFilePath,
			BukuNikahFilePath:      birthDocument.BukuNikahFilePath,
			KkFilePath:             birthDocument.KkFilePath,
			KtpPelaporFilePath:     birthDocument.KtpPelaporFilePath,
			KtpSaksi1FilePath:      birthDocument.KtpSaksi1FilePath,
			KtpSaksi2FilePath:      birthDocument.KtpSaksi2FilePath,
			KtpOrangTuaFilePath:    birthDocument.KtpOrangTuaFilePath,
		})
	}

	response.BirthDocumentResponses = birthDocumentResponses

	// Mengambil semua data dari tabel death_certificate_documents dengan status "belum"
	var deathDocuments []entity.DeathCertificateDocument
	err = tx.Table("death_certificate_documents").Joins("User").Joins("StatementTypeItem.StatementType").Where("status = ?", "belum").Find(&deathDocuments).Error
	if err != nil {
		return err
	}

	var deathDocumentResponses []model.DeathCertificateDocumentWithUserResponse
	for _, deathDocument := range deathDocuments {
		deathDocumentResponses = append(deathDocumentResponses, model.DeathCertificateDocumentWithUserResponse{
			ID:                    deathDocument.ID,
			StatementTypeItem:     deathDocument.StatementTypeItem.Name,
			StatementType:         deathDocument.StatementTypeItem.StatementType.Name,
			Username:              deathDocument.User.Username,
			Name:                  deathDocument.User.Name,
			RtRwFilePath:          deathDocument.RtRwFilePath,
			FormulirFilePath:      deathDocument.FormulirFilePath,
			SuratKematianFilePath: deathDocument.SuratKematianFilePath,
			KtpFilePath:           deathDocument.KtpFilePath,
			KkFilePath:            deathDocument.KkFilePath,
			KtpPelaporFilePath:    deathDocument.KtpPelaporFilePath,
			KtpSaksi1FilePath:     deathDocument.KtpSaksi1FilePath,
			KtpSaksi2FilePath:     deathDocument.KtpSaksi2FilePath,
			BukuNikahFilePath:     deathDocument.BukuNikahFilePath,
		})
	}

	response.DeathDocumentResponses = deathDocumentResponses

	// Mengambil semua data dari tabel marriage_statement_documents dengan status "belum"
	var marriageDocuments []entity.MarriageStatementDocument
	err = tx.Table("marriage_statement_documents").Joins("User").Joins("StatementTypeItem.StatementType").Where("status = ?", "belum").Find(&marriageDocuments).Error
	if err != nil {
		return err
	}

	var marriageDocumentsResponses []model.MarriageStatementDocumentWithUserResponse
	for _, marriageDocument := range marriageDocuments {
		marriageDocumentsResponses = append(marriageDocumentsResponses, model.MarriageStatementDocumentWithUserResponse{
			ID:                marriageDocument.ID,
			StatementTypeItem: marriageDocument.StatementTypeItem.Name,
			StatementType:     marriageDocument.StatementTypeItem.StatementType.Name,
			Username:          marriageDocument.User.Username,
			Name:              marriageDocument.User.Name,
			KkFilePath:        marriageDocument.KkFilePath,
			KtpFilePath:       marriageDocument.KtpFilePath,
			IjazahFilePath:    marriageDocument.IjazahFilePath,
			AktaFilePath:      marriageDocument.AktaFilePath,
			KtpSaksiFilePath:  marriageDocument.KtpSaksiFilePath,
		})
	}

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
