package test

import (
	"one-file/models"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {

	testing.Init()
	os.Setenv("testing", "true")

	models.Build()

	gin.SetMode(gin.TestMode)
}
