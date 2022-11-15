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

func TestFileGet(t *testing.T) {

	createUserInput := controllers.CreateUserInput{
		Username: "another_dummy_user_with_a_file",
		Password: "another_dummy_user",
	}

	// create a user with a linked file
	createUserRequest(createUserInput)

	// get the created user from the db
	user := models.User{}
	models.DB().Last(&user)

	mockToken, _ := auth.CreateToken(user.ID)
	mockResponse := ``

	router := gin.Default()
	router.GET("/file", controllers.GetFile)

	req, _ := http.NewRequest("GET", "/file", &bytes.Buffer{})
	req.Header.Add("Authorization", "Bearer "+string(mockToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, string(responseData), mockResponse)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestFileMod(t *testing.T) {

	createUserInput := controllers.CreateUserInput{
		Username: "another_dummy_user_that_edit_a_file",
		Password: "another_dummy_user",
	}

	// create a user with a linked file
	createUserRequest(createUserInput)

	// get the created user from the db
	user := models.User{}
	models.DB().Last(&user)

	mockToken, _ := auth.CreateToken(user.ID)
	mockResponse := `content`

	router := gin.Default()
	router.GET("/file", controllers.GetFile)
	router.PUT("/file", controllers.ModifyFile)

	// modify the empty file
	json, _ := json.Marshal(controllers.ModifyFileInput{
		Content: "test_content",
	})

	req, _ := http.NewRequest("PUT", "/file", bytes.NewBuffer(json))
	req.Header.Add("Authorization", "Bearer "+string(mockToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// get the modified file

	req2, _ := http.NewRequest("GET", "/file", &bytes.Buffer{})
	req.Header.Add("Authorization", "Bearer "+string(mockToken))

	// flush output
	ioutil.ReadAll(w.Body)

	w2 := httptest.NewRecorder()
	router.ServeHTTP(w, req2)

	responseData, _ := ioutil.ReadAll(w.Body)

	assert.Equal(t, string(responseData), mockResponse)
	assert.Equal(t, http.StatusForbidden, w2.Code)
}
