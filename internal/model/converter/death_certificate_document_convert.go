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

func DeathDocumentCompliteResponses(deathDocumentResponses *[]model.DeathCertificateDocumentWithUserResponse, deathDocuments *[]entity.DeathCertificateDocument) {
	for _, deathDocument := range *deathDocuments {
		*deathDocumentResponses = append(*deathDocumentResponses, model.DeathCertificateDocumentWithUserResponse{
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
			RtRw:                  deathDocument.User.RwList.NameRw,
		})
	}
}

func DeathDocumentIncompliteResponses(deathDocumentResponses *[]model.DeathCertificateDocumentWithUserResponse, deathDocuments *[]entity.DeathCertificateDocument) {
	for _, deathDocument := range *deathDocuments {
		*deathDocumentResponses = append(*deathDocumentResponses, model.DeathCertificateDocumentWithUserResponse{
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
			RtRw:                  deathDocument.User.RwList.NameRw,
		})
	}
}

func DeathDocumentForUserResponses(deathDocumentResponses *[]model.AllMailItemForUserResponse, deathDocuments *[]entity.DeathCertificateDocument) {
	for _, deathDocument := range *deathDocuments {
		*deathDocumentResponses = append(*deathDocumentResponses, model.AllMailItemForUserResponse{
			ID:                deathDocument.ID,
			StatementTypeItem: deathDocument.StatementTypeItem.Name,
			StatementType:     deathDocument.StatementTypeItem.StatementType.Name,
			Date:              deathDocument.CreatedAt.Format("2006-01-02"),
			RW:                deathDocument.User.RwList.NameRw,
			Status:            deathDocument.Status,
			SuratBalasan:      deathDocument.SuratBalasanFilePath,
		})
	}
}
