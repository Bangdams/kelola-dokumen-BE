package model

type StatementTypeResponse struct {
	ID   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type StatementTypeRequest struct {
	Name string `json:"name" validate:"required"`
}
