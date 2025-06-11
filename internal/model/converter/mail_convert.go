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
