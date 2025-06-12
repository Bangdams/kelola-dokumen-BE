package entity

import "time"

type Mail struct {
	ID                   uint   `gorm:"primaryKey"`
	StatementTypeItemId  uint   `gorm:"not null"`
	UserId               uint   `gorm:"not null"`
	KtpFilePath          string `gorm:"not null"`
	KkFilePath           string `gorm:"not null"`
	RtRwFilePath         string `gorm:"not null"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	SuratBalasanFilePath string            `gorm:"not null"`
	Status               string            `gorm:"not null"`
	User                 User              `gorm:"foreignKey:user_id;references:id"`
	StatementTypeItem    StatementTypeItem `gorm:"foreignKey:statement_type_item_id;references:id"`
}
