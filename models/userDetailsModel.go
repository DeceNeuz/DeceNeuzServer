package models

import (
	"time"

	"gorm.io/gorm"
)

type UserDetails struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
	UserName     string    `gorm:"not null"`
	Posts     []uint     `gorm:"type:integer[]"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}