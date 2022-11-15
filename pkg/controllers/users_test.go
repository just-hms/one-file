package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"one-file/pkg/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestLoginHandler(t *testing.T) {

	models.Build()
	gin.SetMode(gin.ReleaseMode)

	// create an user
	user := models.User{
		Username: "dummy",
		Password: "dummy_password",
	}

	models.DB().Create(&user)

	correctMockRegex := `{"token":".*"}`

	// authorized test

	correctResponseData, wCorrect := loginRequest(LoginInput{
		Username: user.Username,
		Password: user.Password,
	})

	assert.MatchRegex(t, correctResponseData, correctMockRegex)
	assert.Equal(t, http.StatusOK, wCorrect.Code)

	// un authorized test

	wrongMockResponse := `{"error":"record not found"}`
	wrongResponseData, wWrong := loginRequest(LoginInput{
		Username: "wrong",
		Password: user.Password,
	})

	assert.Equal(t, wrongResponseData, wrongMockResponse)
	assert.Equal(t, http.StatusUnauthorized, wWrong.Code)

}

// does a login request with the provided input
// returns the response data
func loginRequest(loginInput LoginInput) (string, *httptest.ResponseRecorder) {
	router := gin.Default()
	router.POST("/login", Login)

	json, _ := json.Marshal(loginInput)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(json))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	return string(responseData), w
}

func TestCreateUser(t *testing.T) {

	models.Build()
	gin.SetMode(gin.ReleaseMode)

	createUserInput := CreateUserInput{
		Username: "another_dummy_user",
		Password: "another_dummy_user",
	}

	correctMockResponse := `{}`
	correctResponseData, wCorrect := createUserRequest(createUserInput)

	assert.Equal(t, correctResponseData, correctMockResponse)
	assert.Equal(t, http.StatusCreated, wCorrect.Code)

	// second request

	wrongMockResponse := `{"error":"Username already used"}`
	wrongResponseData, wWrong := createUserRequest(createUserInput)

	assert.Equal(t, wrongResponseData, wrongMockResponse)
	assert.Equal(t, http.StatusForbidden, wWrong.Code)

}

// does a login request with the provided input
// returns the response data
func createUserRequest(createUserInput CreateUserInput) (string, *httptest.ResponseRecorder) {

	router := gin.Default()
	router.POST("/user", CreateUser)

	json, _ := json.Marshal(createUserInput)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(json))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	return string(responseData), w
}
