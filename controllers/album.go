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

type UpdateAlbumInput struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
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

func FindAlbumById(c *gin.Context) {
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var input UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAlbum := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}

	models.DB.Model(&album).Updates(&updatedAlbum)
	c.JSON(http.StatusOK, gin.H{"data": album})
}

func DeleteAlbum(c *gin.Context) {
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&album)
	c.JSON(http.StatusNoContent, gin.H{"deleted": true})
}
