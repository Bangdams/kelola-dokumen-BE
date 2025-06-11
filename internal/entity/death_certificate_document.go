package entity

import "time"

type DeathCertificateDocument struct {
	ID                    uint   `gorm:"primaryKey"`
	StatementTypeItemId   uint   `gorm:"not null"`
	UserId                uint   `gorm:"not null"`
	RtRwFilePath          string `gorm:"not null"`
	FormulirFilePath      string `gorm:"not null"`
	SuratKematianFilePath string `gorm:"not null"`
	KtpFilePath           string `gorm:"not null"`
	KkFilePath            string `gorm:"not null"`
	KtpPelaporFilePath    string `gorm:"not null"`
	KtpSaksi1FilePath     string `gorm:"not null"`
	KtpSaksi2FilePath     string `gorm:"not null"`
	BukuNikahFilePath     string `gorm:"not null"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	User                  User              `gorm:"foreignKey:user_id;references:id"`
	StatementTypeItem     StatementTypeItem `gorm:"foreignKey:statement_type_item_id;references:id"`
}
