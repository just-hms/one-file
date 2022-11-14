package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"unique"`
	Password string

	Name     string
	LastName string

	Birthday time.Time
	LastSeen time.Time

	IsAdmin bool `gorm:"default:false"`
}
