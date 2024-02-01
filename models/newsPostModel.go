package models

import (
	"time"
	"gorm.io/gorm"
)

type NewsPost struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	ShortDescription   string    `gorm:"not null"`
	User     string    `gorm:"not null"`
	Likes     uint      `gorm:"default:0"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}