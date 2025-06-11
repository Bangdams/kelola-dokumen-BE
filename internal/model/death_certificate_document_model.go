package model

type DeathCertificateDocumentResponse struct {
	ID                    uint   `json:"id" validate:"required"`
	StatementTypeItemId   uint   `json:"statement_type_item_id" validate:"required"`
	UserId                uint   `json:"user_id" validate:"required"`
	RtRwFilePath          string `json:"rt_rw_file_path" validate:"required"`
	FormulirFilePath      string `json:"formulir_file_path" validate:"required"`
	SuratKematianFilePath string `json:"surat_kematian_file_path" validate:"required"`
	KtpFilePath           string `json:"ktp_file_path" validate:"required"`
	KkFilePath            string `json:"kk_file_path" validate:"required"`
	KtpPelaporFilePath    string `json:"ktp_pelapor_file_path" validate:"required"`
	KtpSaksi1FilePath     string `json:"ktp_saksi_1_file_path" validate:"required"`
	KtpSaksi2FilePath     string `json:"ktp_saksi_2_file_path" validate:"required"`
	BukuNikahFilePath     string `json:"buku_nikah_file_path" validate:"required"`
}

type DeathCertificateDocumentRequest struct {
	StatementTypeItemId   uint   `json:"statement_type_item_id" validate:"required"`
	UserId                uint   `json:"user_id" validate:"required"`
	RtRwFilePath          string `json:"rt_rw_file_path" validate:"required"`
	FormulirFilePath      string `json:"formulir_file_path" validate:"required"`
	SuratKematianFilePath string `json:"surat_kematian_file_path" validate:"required"`
	KtpFilePath           string `json:"ktp_file_path" validate:"required"`
	KkFilePath            string `json:"kk_file_path" validate:"required"`
	KtpPelaporFilePath    string `json:"ktp_pelapor_file_path" validate:"required"`
	KtpSaksi1FilePath     string `json:"ktp_saksi_1_file_path" validate:"required"`
	KtpSaksi2FilePath     string `json:"ktp_saksi_2_file_path" validate:"required"`
	BukuNikahFilePath     string `json:"buku_nikah_file_path" validate:"required"`
}

type UpdateDeathCertificateDocumentRequest struct {
	ID                    uint   `json:"id" validate:"required"`
	StatementTypeItemId   uint   `json:"statement_type_item_id" validate:"required"`
	UserId                uint   `json:"user_id" validate:"required"`
	RtRwFilePath          string `json:"rt_rw_file_path" validate:"required"`
	FormulirFilePath      string `json:"formulir_file_path" validate:"required"`
	SuratKematianFilePath string `json:"surat_kematian_file_path" validate:"required"`
	KtpFilePath           string `json:"ktp_file_path" validate:"required"`
	KkFilePath            string `json:"kk_file_path" validate:"required"`
	KtpPelaporFilePath    string `json:"ktp_pelapor_file_path" validate:"required"`
	KtpSaksi1FilePath     string `json:"ktp_saksi_1_file_path" validate:"required"`
	KtpSaksi2FilePath     string `json:"ktp_saksi_2_file_path" validate:"required"`
	BukuNikahFilePath     string `json:"buku_nikah_file_path" validate:"required"`
}
