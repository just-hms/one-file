package controllers

import (
	"net/http"
	"one-file/auth"
	"one-file/models"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	input := LoginInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	err := models.DB().Where("username = ?", input.Username).Take(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	auth.Verify(user.Password, input.Password)

	token, err := auth.CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

type CreateInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Create(c *gin.Context) {

	input := CreateInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB().Where("email = ?", input.Username).First(&models.User{}).Error; err == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Username already used"})
		return
	}

	password, err := auth.HashAndSalt(input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: password,
	}

	models.DB().Create(&user)

	c.JSON(http.StatusCreated, gin.H{})

}
