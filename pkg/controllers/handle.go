package controllers

import (
	"net/http"

	"one-file/pkg/middlewares"

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

	files.GET("", middlewares.RequireAuth, GetFile)
	files.PUT("", middlewares.RequireAuth, ModifyFile)

	router.POST("/user", middlewares.RequireAdmin, CreateUser)

	router.Run()
}

func Base(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Oh no! You found me!"})
	return
}
