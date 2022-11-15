package controllers

import (
	"net/http"
	"one-file/auth"
	"one-file/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenIDFromRequest(c *gin.Context) (uint, error) {
	authHeader := c.GetHeader("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")

	return auth.ExtractTokenID(token)
}

func RequireAuth(c *gin.Context) {

	if _, err := extractTokenIDFromRequest(c); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Next()
}

func RequireAdmin(c *gin.Context) {

	userID, err := extractTokenIDFromRequest(c)

	// check fo any errors
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user = models.User{}
	err = models.DB().Where("id = ? AND is_admin = TRUE", userID).First(&user).Error

	if err != nil || !user.IsAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Next()
}
