package controllers

import (
	"net/http"
	"one-file/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetFile(c *gin.Context) {

	// get the user id from the token

	var (
		user_id uint
		err     error
	)

	if user_id, err = extractTokenIDFromRequest(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find the file associated to the user

	file := models.File{}

	if err := models.DB().Where("user_id = ?", user_id).Take(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file": file.Content})

}

type ModifyFileInput struct {
	Content string `json:"content" binding:"required"`
}

func ModifyFile(c *gin.Context) {

	// get and check the input
	input := ModifyFileInput{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var (
		user_id uint
		err     error
	)

	// get the user id from the token

	if user_id, err = extractTokenIDFromRequest(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find the file associated to the user

	file := models.File{}

	if err := models.DB().Where("user_id = ?", user_id).Take(&file).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	file.Content = input.Content

	models.DB().Save(&file)
}
