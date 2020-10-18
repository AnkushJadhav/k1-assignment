package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null" validate:"required"`
	Email     string    `gorm:"not null,unique" validate:"required"`
	Password  string    `gorm:"not null" validate:"required"`
	Hits      uint      `gorm:"not null,default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
