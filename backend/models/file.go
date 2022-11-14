package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model

	Content string

	// belongs to the user
	UserID int
	User   User
}
