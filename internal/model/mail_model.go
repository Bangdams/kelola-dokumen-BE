package model

type MailResponse struct {
	ID                  uint   `json:"id" validate:"required"`
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	RtRwFilePath        string `json:"rt_rw_file_path" validate:"required"`
}

type MailWithUserResponse struct {
	ID                uint   `json:"id" validate:"required"`
	StatementTypeItem string `json:"statement_type_item" validate:"required"`
	StatementType     string `json:"statement_type" validate:"required"`
	Username          string `json:"username" validate:"required"`
	Name              string `json:"name" validate:"required"`
	KtpFilePath       string `json:"ktp_file_path" validate:"required"`
	KkFilePath        string `json:"kk_file_path" validate:"required"`
	RtRwFilePath      string `json:"rt_rw_file_path" validate:"required"`
	SuratBalasan      string `json:"surat_balasan_file_path" validate:"required"`
}

type MailRequest struct {
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	RtRwFilePath        string `json:"rt_rw_file_path" validate:"required"`
}

type BalasanRequest struct {
	ID                   uint   `json:"id" validate:"required"`
	SuratBalasanFilePath string `json:"surat_balasan" validate:"required"`
}

type UpdateMailRequest struct {
	ID                  uint   `json:"id" validate:"required"`
	StatementTypeItemId uint   `json:"statement_type_item_id" validate:"required"`
	UserId              uint   `json:"user_id" validate:"required"`
	KtpFilePath         string `json:"ktp_file_path" validate:"required"`
	KkFilePath          string `json:"kk_file_path" validate:"required"`
	RtRwFilePath        string `json:"rt_rw_file_path" validate:"required"`
}

type AllMailResponse struct {
	MailResponses             []MailWithUserResponse                      `json:"mails"`
	BirthDocumentResponses    []BirthCertificateDocumentWithUserResponse  `json:"birth_documents"`
	DeathDocumentResponses    []DeathCertificateDocumentWithUserResponse  `json:"death_documents"`
	MarriageDocumentResponses []MarriageStatementDocumentWithUserResponse `json:"marriage_documents"`
}
