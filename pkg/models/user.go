package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"unique"`
	Password string

	File File `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	IsAdmin bool `gorm:"default:false"`
}
