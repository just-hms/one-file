package controllers

import (
	"net/http"
	"one-file/pkg/auth"
	"one-file/pkg/models"

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

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {

	input := CreateUserInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB().Where("username = ?", input.Username).First(&models.User{}).Error; err == nil {
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

	// create an empty file
	file := models.File{
		Content: "",
	}
	models.DB().Save(&file)

	// link it to the new user
	models.DB().Model(&user).Association("File").Append(&file)

	c.JSON(http.StatusCreated, gin.H{})

}
