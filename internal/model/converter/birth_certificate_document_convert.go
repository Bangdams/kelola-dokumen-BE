package converter

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
)

func BirthCertificateDocumentToResponse(birthDocument *entity.BirthCertificateDocument) *model.BirthCertificateDocumentResponse {
	log.Println("log from BirthCertificateDocument to response")

	return &model.BirthCertificateDocumentResponse{
		ID:                     birthDocument.ID,
		StatementTypeItemId:    birthDocument.StatementTypeItemId,
		UserId:                 birthDocument.UserId,
		RtRwFilePath:           birthDocument.RtRwFilePath,
		FormulirFilePath:       birthDocument.FormulirFilePath,
		SuratKelahiranFilePath: birthDocument.SuratKelahiranFilePath,
		BukuNikahFilePath:      birthDocument.BukuNikahFilePath,
		KkFilePath:             birthDocument.KkFilePath,
		KtpPelaporFilePath:     birthDocument.KtpPelaporFilePath,
		KtpSaksi1FilePath:      birthDocument.KtpSaksi1FilePath,
		KtpSaksi2FilePath:      birthDocument.KtpSaksi2FilePath,
		KtpOrangTuaFilePath:    birthDocument.KtpOrangTuaFilePath,
	}
}

func BirthCertificateDocumentToResponses(birthDocuments *[]entity.BirthCertificateDocument) *[]model.BirthCertificateDocumentResponse {
	var birthDocumentResponses []model.BirthCertificateDocumentResponse

	log.Println("log from BirthCertificateDocument to responses")

	for _, birthDocument := range *birthDocuments {
		birthDocumentResponses = append(birthDocumentResponses, *BirthCertificateDocumentToResponse(&birthDocument))
	}

	return &birthDocumentResponses
}

func BirthDocumentCompliteResponses(birthDocumentResponses *[]model.BirthCertificateDocumentWithUserResponse, birthDocuments *[]entity.BirthCertificateDocument) {
	for _, birthDocument := range *birthDocuments {
		*birthDocumentResponses = append(*birthDocumentResponses, model.BirthCertificateDocumentWithUserResponse{
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
			RtRw:                   birthDocument.User.RwList.NameRw,
		})
	}
}

func BirthDocumentIncompliteResponses(birthDocumentResponses *[]model.BirthCertificateDocumentWithUserResponse, birthDocuments *[]entity.BirthCertificateDocument) {
	for _, birthDocument := range *birthDocuments {
		*birthDocumentResponses = append(*birthDocumentResponses, model.BirthCertificateDocumentWithUserResponse{
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
			RtRw:                   birthDocument.User.RwList.NameRw,
		})
	}
}

func BirthDocumentForUserResponses(birthDocumentResponses *[]model.AllMailItemForUserResponse, birthDocuments *[]entity.BirthCertificateDocument) {
	for _, birthDocument := range *birthDocuments {
		*birthDocumentResponses = append(*birthDocumentResponses, model.AllMailItemForUserResponse{
			ID:                birthDocument.ID,
			StatementTypeItem: birthDocument.StatementTypeItem.Name,
			StatementType:     birthDocument.StatementTypeItem.StatementType.Name,
			Date:              birthDocument.CreatedAt.Format("2006-01-02"),
			RW:                birthDocument.User.RwList.NameRw,
			Status:            birthDocument.Status,
			SuratBalasan:      birthDocument.SuratBalasanFilePath,
		})
	}
}
