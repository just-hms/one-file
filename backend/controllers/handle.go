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

	users := router.Group("/users")
	users.GET("/login", Login)

	router.Run()
}

func Base(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Oh no! You found me!"})
	return
}
