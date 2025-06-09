package entity

import "time"

type RwList struct {
	ID        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"not null"`
	NameRw    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      *User `gorm:"foreignKey:user_id;references:id"`
}
