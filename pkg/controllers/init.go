package controllers

import (
	"one-file/pkg/models"
	"os"

	"github.com/gin-gonic/gin"
)

func initTest() {
	os.Setenv("testing", "true")
	models.Build()
	gin.SetMode(gin.TestMode)
}
