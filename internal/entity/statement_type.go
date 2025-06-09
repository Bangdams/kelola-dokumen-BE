package entity

import "time"

type StatementType struct {
	ID                uint `gorm:"primaryKey"`
	Name              uint `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	StatementTypeItem []StatementTypeItem `gorm:"foreignKey:statement_type_id;references:id"`
}
