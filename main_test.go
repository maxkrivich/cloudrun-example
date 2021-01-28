package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestPing(t *testing.T) {
	body := gin.H{
		"message": "pong",
	}
	router := SetupRouter()
	w := performRequest(router, "GET", "/api/ping")

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["message"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["message"], value)
}

func TestHelloAnonymous(t *testing.T) {
	body := gin.H{
		"message": "Hello, Anonymous!",
	}
	router := SetupRouter()
	w := performRequest(router, "GET", "/api/hello")

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["message"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["message"], value)
}

func TestHelloWithName(t *testing.T) {
	body := gin.H{
		"message": "Hello, Max!",
	}
	router := SetupRouter()
	w := performRequest(router, "GET", "/api/hello?name=Max")

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["message"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["message"], value)
}
