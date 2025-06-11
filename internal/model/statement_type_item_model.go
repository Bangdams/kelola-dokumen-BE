package model

type StatementTypeItemResponse struct {
	ID              uint   `json:"id" validate:"required"`
	StatementTypeId uint   `json:"statement_type_id" validate:"required"`
	Name            string `json:"name" validate:"required"`
}

type StatementTypeItemRequest struct {
	StatementTypeId uint   `json:"statement_type_id" validate:"required"`
	Name            string `json:"name" validate:"required"`
}

type UpdateStatementTypeItemRequest struct {
	ID              uint   `json:"id" validate:"required"`
	StatementTypeId uint   `json:"statement_type_id" validate:"required"`
	Name            string `json:"name" validate:"required"`
}
