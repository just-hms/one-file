package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model

	Content string
	UserID  uint
}
