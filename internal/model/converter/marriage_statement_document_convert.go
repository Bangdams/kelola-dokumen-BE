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

func MarriageDocumentsCompliteResponses(marriageDocumentsResponses *[]model.MarriageStatementDocumentWithUserResponse, marriageDocuments *[]entity.MarriageStatementDocument) {
	for _, marriageDocument := range *marriageDocuments {
		*marriageDocumentsResponses = append(*marriageDocumentsResponses, model.MarriageStatementDocumentWithUserResponse{
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
			RtRw:              marriageDocument.User.RwList.NameRw,
		})
	}

}
func MarriageDocumentsIncompliteResponses(marriageDocumentsResponses *[]model.MarriageStatementDocumentWithUserResponse, marriageDocuments *[]entity.MarriageStatementDocument) {
	for _, marriageDocument := range *marriageDocuments {
		*marriageDocumentsResponses = append(*marriageDocumentsResponses, model.MarriageStatementDocumentWithUserResponse{
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
			RtRw:              marriageDocument.User.RwList.NameRw,
		})
	}
}

func MarriageDocumentForUserResponses(marriageDocumentsResponses *[]model.AllMailItemForUserResponse, marriageDocuments *[]entity.MarriageStatementDocument) {
	for _, marriageDocument := range *marriageDocuments {
		*marriageDocumentsResponses = append(*marriageDocumentsResponses, model.AllMailItemForUserResponse{
			ID:                marriageDocument.ID,
			StatementTypeItem: marriageDocument.StatementTypeItem.Name,
			StatementType:     marriageDocument.StatementTypeItem.StatementType.Name,
			Date:              marriageDocument.CreatedAt.Format("2006-01-02"),
			RW:                marriageDocument.User.RwList.NameRw,
			Status:            marriageDocument.Status,
			SuratBalasan:      marriageDocument.SuratBalasanFilePath,
		})
	}
}
