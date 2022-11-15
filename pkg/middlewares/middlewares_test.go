package middlewares

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"one-file/pkg/auth"
	"one-file/pkg/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestAuthMiddleWare(t *testing.T) {

	models.Build()
	gin.SetMode(gin.ReleaseMode)

	user := models.User{
		Username: "dummy_user",
		Password: "dummy_password",
	}

	models.DB().Save(&user)
	// authorized test
	correctMockToken, _ := auth.CreateToken(user.ID)
	correctMockResponse := `{"res":"ok"}`

	correctResponseData, wCorrect := fakeRouter(correctMockToken, RequireAuth)

	assert.Equal(t, correctResponseData, correctMockResponse)
	assert.Equal(t, http.StatusOK, wCorrect.Code)

	// unauthorized test

	wrongMockToken := "something_wrong"
	wrongMockResponse := `{"error":"Unauthorized"}`

	wrongResponseData, wWrong := fakeRouter(wrongMockToken, RequireAuth)

	assert.Equal(t, wrongResponseData, wrongMockResponse)
	assert.Equal(t, http.StatusUnauthorized, wWrong.Code)
}

func TestAdminMiddleWare(t *testing.T) {

	models.Build()
	gin.SetMode(gin.ReleaseMode)

	admin := models.User{
		Username: "dummy_admin",
		IsAdmin:  true,
		Password: "dummy_password",
	}

	models.DB().Save(&admin)
	// authorized test
	correctMockToken, _ := auth.CreateToken(admin.ID)
	correctMockResponse := `{"res":"ok"}`

	correctResponseData, wCorrect := fakeRouter(correctMockToken, RequireAdmin)

	assert.Equal(t, correctResponseData, correctMockResponse)
	assert.Equal(t, http.StatusOK, wCorrect.Code)

	// unauthorized test

	wrongMockToken := "something_wrong"
	wrongMockResponse := `{"error":"Unauthorized"}`

	wrongResponseData, wWrong := fakeRouter(wrongMockToken, RequireAdmin)

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
