package entity

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Name      string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	RwList    RwList `gorm:"foreignKey:user_id;references:id"`
	Mail      []Mail `gorm:"foreignKey:user_id;references:id"`
}
