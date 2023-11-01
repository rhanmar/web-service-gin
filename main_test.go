package main

import (
	"bytes"
	"encoding/json"
	"example/web-service-gin/pkg/config"
	"example/web-service-gin/pkg/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindAlbums(t *testing.T) {
	router := SetUpRouter()
	config.ConnectTestDatabase()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockResponse := `{"data":[]}`
	responseBody, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseBody))

	config.ClearTestDatabase()
}

func TestCreateAlbum(t *testing.T) {
	router := SetUpRouter()
	config.ConnectTestDatabase()

	w := httptest.NewRecorder()
	newAlbum := models.Album{
		Title:  "test album",
		Artist: "test artist",
		Price:  1.1,
	}
	body, _ := json.Marshal(newAlbum)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	requestBody, _ := ioutil.ReadAll(w.Body)
	var dat map[string]interface{}
	json.Unmarshal(requestBody, &dat)
	assert.Equal(t, true, dat["created"])
	_, ok := dat["data"]
	assert.Equal(t, true, ok)

	config.ClearTestDatabase()
}
