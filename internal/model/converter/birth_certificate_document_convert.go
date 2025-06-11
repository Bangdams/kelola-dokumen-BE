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
