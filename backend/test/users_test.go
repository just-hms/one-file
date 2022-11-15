package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"one-file/auth"
	"one-file/controllers"
	"one-file/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestLoginHandler(t *testing.T) {

	// create an user
	user := models.User{
		Username: "dummy",
		Password: "dummy_password",
	}

	models.DB().Create(&user)

	correctMockRegex := `{"token":".*"}`

	// authorized test

	correctResponseData, wCorrect := loginRequest(controllers.LoginInput{
		Username: user.Username,
		Password: user.Password,
	})

	assert.MatchRegex(t, correctResponseData, correctMockRegex)
	assert.Equal(t, http.StatusOK, wCorrect.Code)

	// un authorized test

	wrongMockResponse := `{"error":"record not found"}`
	wrongResponseData, wWrong := loginRequest(controllers.LoginInput{
		Username: "wrong",
		Password: user.Password,
	})

	assert.Equal(t, wrongResponseData, wrongMockResponse)
	assert.Equal(t, http.StatusUnauthorized, wWrong.Code)

}

// does a login request with the provided input
// returns the response data
func loginRequest(loginInput controllers.LoginInput) (string, *httptest.ResponseRecorder) {
	router := gin.Default()
	router.POST("/login", controllers.Login)

	json, _ := json.Marshal(loginInput)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(json))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	return string(responseData), w
}

func TestAuthMiddleWare(t *testing.T) {

	user := models.User{
		Username: "dummy_user",
		Password: "dummy_password",
	}

	models.DB().Save(&user)
	// authorized test
	correctMockToken, _ := auth.CreateToken(user.ID)
	correctMockResponse := `{"res":"ok"}`

	correctResponseData, wCorrect := fakeRouter(correctMockToken, controllers.RequireAuth)

	assert.Equal(t, correctResponseData, correctMockResponse)
	assert.Equal(t, http.StatusOK, wCorrect.Code)

	// unauthorized test

	wrongMockToken := "something_wrong"
	wrongMockResponse := `{"error":"Unauthorized"}`

	wrongResponseData, wWrong := fakeRouter(wrongMockToken, controllers.RequireAuth)

	assert.Equal(t, wrongResponseData, wrongMockResponse)
	assert.Equal(t, http.StatusUnauthorized, wWrong.Code)
}

func TestAdminMiddleWare(t *testing.T) {

	admin := models.User{
		Username: "dummy_admin",
		IsAdmin:  true,
		Password: "dummy_password",
	}

	models.DB().Save(&admin)
	// authorized test
	correctMockToken, _ := auth.CreateToken(admin.ID)
	correctMockResponse := `{"res":"ok"}`

	correctResponseData, wCorrect := fakeRouter(correctMockToken, controllers.RequireAdmin)

	assert.Equal(t, correctResponseData, correctMockResponse)
	assert.Equal(t, http.StatusOK, wCorrect.Code)

	// unauthorized test

	wrongMockToken := "something_wrong"
	wrongMockResponse := `{"error":"Unauthorized"}`

	wrongResponseData, wWrong := fakeRouter(wrongMockToken, controllers.RequireAdmin)

	assert.Equal(t, wrongResponseData, wrongMockResponse)
	assert.Equal(t, http.StatusUnauthorized, wWrong.Code)
}

// given a token and a middleware, fakes a router and
// returns the response and the response status
func fakeRouter(mockToken string, middeware func(c *gin.Context)) (string, *httptest.ResponseRecorder) {
	router := gin.Default()

	router.GET("/auth", middeware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"res": "ok"})
	})

	req, _ := http.NewRequest("GET", "/auth", &bytes.Buffer{})
	req.Header.Add("Authorization", "Bearer "+string(mockToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	return string(responseData), w

}

func TestCreateUser(t *testing.T) {

	createUserInput := controllers.CreateUserInput{
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
func createUserRequest(createUserInput controllers.CreateUserInput) (string, *httptest.ResponseRecorder) {

	router := gin.Default()
	router.POST("/user", controllers.CreateUser)

	json, _ := json.Marshal(createUserInput)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(json))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	return string(responseData), w
}
