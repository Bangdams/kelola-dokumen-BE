package converter

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
)

func StatementTypeToResponse(statementType *entity.StatementType) *model.StatementTypeResponse {
	log.Println("log from StatementType to response")

	return &model.StatementTypeResponse{
		ID:   statementType.ID,
		Name: statementType.Name,
	}
}

func StatementTypeToResponses(statementTypeItems *[]entity.StatementType) *[]model.StatementTypeResponse {
	var statementTypeResponses []model.StatementTypeResponse

	log.Println("log from StatementType to responses")

	for _, statementTypeItem := range *statementTypeItems {
		statementTypeResponses = append(statementTypeResponses, *StatementTypeToResponse(&statementTypeItem))
	}

	return &statementTypeResponses
}
