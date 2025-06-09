package entity

import "time"

type StatementTypeItem struct {
	ID              uint   `gorm:"primaryKey"`
	StatementTypeId uint   `gorm:"not null"`
	Name            string `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	StatementType   StatementType `gorm:"foreignKey:statement_type_id;references:id"`
	Mail            []Mail        `gorm:"foreignKey:statement_type_item_id;references:id"`
}
