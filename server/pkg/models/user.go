package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null,unique"`
	Password  string `gorm:"not null"`
	Hits      uint   `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool `gorm:"not null,default:true"`
}
