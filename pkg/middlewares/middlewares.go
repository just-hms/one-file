package middlewares

import (
	"net/http"
	"one-file/pkg/auth"
	"one-file/pkg/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractTokenIDFromRequest(c *gin.Context) (uint, error) {
	authHeader := c.GetHeader("Authorization")

	token := strings.TrimPrefix(authHeader, "Bearer ")

	return auth.ExtractTokenID(token)
}

func RequireAuth(c *gin.Context) {

	if _, err := ExtractTokenIDFromRequest(c); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Next()
}

func RequireAdmin(c *gin.Context) {

	userID, err := ExtractTokenIDFromRequest(c)

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
