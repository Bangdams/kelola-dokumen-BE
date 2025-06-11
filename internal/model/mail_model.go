package model

type MailResponse struct {
	ID                  uint   `json:"id" validate:"required"`
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	RtRwFilePath        string `json:"rt_rw_file_path" validate:"required"`
}

type MailRequest struct {
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	RtRwFilePath        string `json:"rt_rw_file_path" validate:"required"`
}

type UpdateMailRequest struct {
	ID                  uint   `json:"id" validate:"required"`
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	RtRwFilePath        string `json:"rt_rw_file_path" validate:"required"`
}
