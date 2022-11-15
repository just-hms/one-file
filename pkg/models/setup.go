package models

import (
	"one-file/internal/constants"
	"one-file/pkg/auth"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB = nil

func connect() (db *gorm.DB) {

	connection_string := ""
	my_logger := logger.Default

	if os.Getenv("testing") == "true" {
		connection_string = "file::memory:?cache=shared"
		my_logger = logger.Discard
	} else {
		connection_string = constants.DATABASE_NAME
	}

	db, err := gorm.Open(sqlite.Open(connection_string), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   my_logger,
	})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// singleton

func DB() (db *gorm.DB) {
	if database == nil {
		database = connect()
	}
	return database
}

func Build() {
	DB().AutoMigrate(
		&User{},
		&File{},
	)

	password, _ := auth.HashAndSalt(constants.ADMIN_PASSWORD)

	DB().Create(&User{
		Username: constants.ADMIN_USERNAME,
		Password: password,
		IsAdmin:  true,
	})

}
