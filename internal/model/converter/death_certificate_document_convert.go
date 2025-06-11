package converter

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
)

func DeathCertificateDocumentToResponse(deathDocument *entity.DeathCertificateDocument) *model.DeathCertificateDocumentResponse {
	log.Println("log from DeathCertificateDocument to response")

	return &model.DeathCertificateDocumentResponse{
		ID:                    deathDocument.ID,
		StatementTypeItemId:   deathDocument.StatementTypeItemId,
		UserId:                deathDocument.UserId,
		RtRwFilePath:          deathDocument.RtRwFilePath,
		FormulirFilePath:      deathDocument.FormulirFilePath,
		SuratKematianFilePath: deathDocument.SuratKematianFilePath,
		KtpFilePath:           deathDocument.KtpFilePath,
		KkFilePath:            deathDocument.KkFilePath,
		KtpPelaporFilePath:    deathDocument.KtpPelaporFilePath,
		KtpSaksi1FilePath:     deathDocument.KtpSaksi1FilePath,
		KtpSaksi2FilePath:     deathDocument.KtpSaksi2FilePath,
		BukuNikahFilePath:     deathDocument.BukuNikahFilePath,
	}
}

func DeathCertificateDocumentToResponses(deathDocuments *[]entity.DeathCertificateDocument) *[]model.DeathCertificateDocumentResponse {
	var deathDocumentResponses []model.DeathCertificateDocumentResponse

	log.Println("log from DeathCertificateDocument to responses")

	for _, deathDocument := range *deathDocuments {
		deathDocumentResponses = append(deathDocumentResponses, *DeathCertificateDocumentToResponse(&deathDocument))
	}

	return &deathDocumentResponses
}
