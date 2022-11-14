package controllers

import (
	"net/http"
	"one-file/auth"
	"one-file/models"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

func Login(c *gin.Context) {

	input := LoginInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	err := models.DB().Model(models.User{}).Where("email = ?", input.Email).Take(&user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth.Verify(user.Password, input.Password)

	token, err := auth.CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": token})

}
