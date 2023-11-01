package controllers

import (
	"example/web-service-gin/pkg/config"
	"example/web-service-gin/pkg/models"
	"example/web-service-gin/pkg/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAlbum(c *gin.Context) {
	var input schemas.CreateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	album := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}
	config.DB.Create(&album)
	c.JSON(http.StatusCreated, gin.H{"data": album, "created": true})
}

func FindAlbums(c *gin.Context) {
	var albums []models.Album
	config.DB.Find(&albums)
	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func FindAlbumById(c *gin.Context) {
	var album models.Album
	if err := config.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	if err := config.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var input schemas.UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAlbum := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}

	config.DB.Model(&album).Updates(&updatedAlbum)
	c.JSON(http.StatusOK, gin.H{"data": album})
}

func DeleteAlbum(c *gin.Context) {
	var album models.Album
	if err := config.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	config.DB.Delete(&album)
	c.JSON(http.StatusNoContent, gin.H{"deleted": true})
}
