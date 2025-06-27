package converter

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
)

func MailToResponse(mail *entity.Mail) *model.MailResponse {
	log.Println("log from Mail to response")

	return &model.MailResponse{
		ID:                  mail.ID,
		StatementTypeItemId: mail.StatementTypeItemId,
		UserId:              mail.UserId,
		KtpFilePath:         mail.KtpFilePath,
		KkFilePath:          mail.KkFilePath,
		RtRwFilePath:        mail.RtRwFilePath,
	}
}

func MailToResponses(mails *[]entity.Mail) *[]model.MailResponse {
	var mailResponses []model.MailResponse

	log.Println("log from Mail to responses")

	for _, mail := range *mails {
		mailResponses = append(mailResponses, *MailToResponse(&mail))
	}

	return &mailResponses
}

func MailCompliteResponses(mailResponses *[]model.MailWithUserResponse, mails *[]entity.Mail) {
	for _, mail := range *mails {
		*mailResponses = append(*mailResponses, model.MailWithUserResponse{
			ID:                mail.ID,
			StatementTypeItem: mail.StatementTypeItem.Name,
			StatementType:     mail.StatementTypeItem.StatementType.Name,
			Username:          mail.User.Username,
			Name:              mail.User.Name,
			KtpFilePath:       mail.KtpFilePath,
			KkFilePath:        mail.KkFilePath,
			RtRwFilePath:      mail.RtRwFilePath,
			SuratBalasan:      mail.SuratBalasanFilePath,
			RtRw:              mail.User.RwList.NameRw,
		})
	}
}

func MailIncompliteResponses(mailResponses *[]model.MailWithUserResponse, mails *[]entity.Mail) {
	for _, mail := range *mails {
		*mailResponses = append(*mailResponses, model.MailWithUserResponse{
			ID:                mail.ID,
			StatementTypeItem: mail.StatementTypeItem.Name,
			StatementType:     mail.StatementTypeItem.StatementType.Name,
			Username:          mail.User.Username,
			Name:              mail.User.Name,
			KtpFilePath:       mail.KtpFilePath,
			KkFilePath:        mail.KkFilePath,
			RtRwFilePath:      mail.RtRwFilePath,
			RtRw:              mail.User.RwList.NameRw,
		})
	}
}

func MailForUserResponses(mailResponses *[]model.AllMailItemForUserResponse, mails *[]entity.Mail) {
	for _, mail := range *mails {
		*mailResponses = append(*mailResponses, model.AllMailItemForUserResponse{
			ID:                mail.ID,
			StatementTypeItem: mail.StatementTypeItem.Name,
			StatementType:     mail.StatementTypeItem.StatementType.Name,
			Date:              mail.CreatedAt.Format("2006-01-02"),
			RW:                mail.User.RwList.NameRw,
			Status:            mail.Status,
			SuratBalasan:      mail.SuratBalasanFilePath,
		})
	}
}
