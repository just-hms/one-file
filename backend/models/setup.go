package models

import (
	"one-file/auth"
	"one-file/constants"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

func connect() (db *gorm.DB) {

	connection_string := ""

	if os.Getenv("testing") == "true" {
		connection_string = "file::memory:?cache=shared"
	} else {
		connection_string = constants.DATABASE_NAME
	}

	db, err := gorm.Open(sqlite.Open(connection_string), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
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
	)

	password, _ := auth.HashAndSalt(constants.ADMIN_PASSWORD)

	DB().Create(&User{
		Username: constants.ADMIN_USERNAME,
		Password: password,
	})

}
