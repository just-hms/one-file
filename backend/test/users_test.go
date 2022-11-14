package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"one-file/auth"
	"one-file/constants"
	"one-file/controllers"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

// test that given a correct input the system returns a token
func TestLoginHandler(t *testing.T) {

	mockRegex := `{"token":".*"}`

	router := gin.Default()

	router.POST("/login", controllers.Login)

	loginInput := controllers.LoginInput{
		Username: constants.ADMIN_USERNAME,
		Password: constants.ADMIN_PASSWORD,
	}

	json, _ := json.Marshal(loginInput)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(json))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.MatchRegex(t, string(responseData), mockRegex)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestAuth(t *testing.T) {

	mockToken, _ := auth.CreateToken(1)

	mockResponse := `{"res":"ok"}`

	router := gin.Default()

	router.GET("/auth", controllers.RequireAuth, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"res": "ok"})
	})

	req, _ := http.NewRequest("GET", "/auth", &bytes.Buffer{})
	req.Header.Add("Authorization", "Bearer "+string(mockToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, string(responseData), mockResponse)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAdmin(t *testing.T) {

	mockToken, _ := auth.CreateToken(1)

	mockResponse := `{"res":"ok"}`

	router := gin.Default()

	router.GET("/auth", controllers.RequireAuth, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"res": "ok"})
	})

	req, _ := http.NewRequest("GET", "/auth", &bytes.Buffer{})
	req.Header.Add("Authorization", "Bearer "+string(mockToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, string(responseData), mockResponse)
	assert.Equal(t, http.StatusOK, w.Code)
}
