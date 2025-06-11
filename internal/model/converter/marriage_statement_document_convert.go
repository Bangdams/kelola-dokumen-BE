package converter

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
)

func MarriageStatementDocumentToResponse(marriageDocument *entity.MarriageStatementDocument) *model.MarriageStatementDocumentResponse {
	log.Println("log from MarriageStatementDocument to response")

	return &model.MarriageStatementDocumentResponse{
		ID:                  marriageDocument.ID,
		StatementTypeItemId: marriageDocument.StatementTypeItemId,
		UserId:              marriageDocument.UserId,
		KkFilePath:          marriageDocument.KkFilePath,
		KtpFilePath:         marriageDocument.KtpFilePath,
		IjazahFilePath:      marriageDocument.IjazahFilePath,
		AktaFilePath:        marriageDocument.AktaFilePath,
		KtpSaksiFilePath:    marriageDocument.KtpSaksiFilePath,
	}
}

func MarriageStatementDocumentToResponses(marriageDocuments *[]entity.MarriageStatementDocument) *[]model.MarriageStatementDocumentResponse {
	var marriageDocumentResponses []model.MarriageStatementDocumentResponse

	log.Println("log from MarriageStatementDocument to responses")

	for _, marriageDocument := range *marriageDocuments {
		marriageDocumentResponses = append(marriageDocumentResponses, *MarriageStatementDocumentToResponse(&marriageDocument))
	}

	return &marriageDocumentResponses
}
