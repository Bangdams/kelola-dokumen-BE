package model

type MarriageStatementDocumentResponse struct {
	ID                  uint   `json:"id" validate:"required"`
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	IjazahFilePath      string `json:"ijazah_file_path" validate:"required"`
	AktaFilePath        string `json:"akta_file_path" validate:"required"`
	KtpSaksiFilePath    string `json:"ktp_saksi_file_path" validate:"required"`
}

type MarriageStatementDocumentRequest struct {
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	IjazahFilePath      string `json:"ijazah_file_path" validate:"required"`
	AktaFilePath        string `json:"akta_file_path" validate:"required"`
	KtpSaksiFilePath    string `json:"ktp_saksi_file_path" validate:"required"`
}

type UpdateMarriageStatementDocumentRequest struct {
	ID                  uint   `json:"id" validate:"required"`
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	IjazahFilePath      string `json:"ijazah_file_path" validate:"required"`
	AktaFilePath        string `json:"akta_file_path" validate:"required"`
	KtpSaksiFilePath    string `json:"ktp_saksi_file_path" validate:"required"`
}
