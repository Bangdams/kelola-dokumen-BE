package converter

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
)

func StatementTypeItemToResponse(statementTypeItem *entity.StatementTypeItem) *model.StatementTypeItemResponse {
	log.Println("log from StatementTypeItem to response")

	return &model.StatementTypeItemResponse{
		ID:              statementTypeItem.ID,
		StatementTypeId: statementTypeItem.StatementTypeId,
		Name:            statementTypeItem.Name,
	}
}

func StatementTypeItemToResponses(statementTypeItems *[]entity.StatementTypeItem) *[]model.StatementTypeItemResponse {
	var statementTypeItemsResponses []model.StatementTypeItemResponse

	log.Println("log from StatementTypeItem to responses")

	for _, statementTypeItem := range *statementTypeItems {
		statementTypeItemsResponses = append(statementTypeItemsResponses, *StatementTypeItemToResponse(&statementTypeItem))
	}

	return &statementTypeItemsResponses
}
