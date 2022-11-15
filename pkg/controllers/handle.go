package controllers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", Base)

	users := router.Group("/user")
	users.GET("/login", Login)

	files := router.Group("/file")

	files.GET("", RequireAuth, GetFile)
	files.PUT("", RequireAuth, ModifyFile)

	router.POST("/user", RequireAdmin, CreateUser)

	router.Run()
}

func Base(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Oh no! You found me!"})
	return
}
