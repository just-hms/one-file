package test

import (
	"one-file/auth"
	"one-file/constants"
	"one-file/models"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {

	testing.Init()
	os.Setenv("testing", "true")

	models.Build()

	// create a dummy user for testing
	password, _ := auth.HashAndSalt(constants.DUMMY_PASSWORD)
	models.DB().Create(&models.User{
		Username: constants.DUMMY_USERNAME,
		Password: password,
		IsAdmin:  false,
	})

	gin.SetMode(gin.TestMode)
}
