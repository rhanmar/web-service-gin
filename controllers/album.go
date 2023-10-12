package controllers

import (
	"example/web-service-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateAlbumInput struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

func CreateAlbum(c *gin.Context) {
	var input CreateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	album := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}
	models.DB.Create(&album)
	c.JSON(http.StatusCreated, gin.H{"data": album, "created": true})
}

func FindAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)
	c.JSON(http.StatusOK, gin.H{"data": albums})
}
