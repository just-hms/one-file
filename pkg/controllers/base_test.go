package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestBaseHandler(t *testing.T) {
	initTest()
	mockResponse := `{"data":"Oh no! You found me!"}`

	router := gin.Default()
	router.GET("/", Base)

	req, _ := http.NewRequest("GET", "/", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, string(responseData), mockResponse)
	assert.Equal(t, w.Code, http.StatusOK)
}
