package controllers

import (
	"example/web-service-gin/pkg/schemas"
	"example/web-service-gin/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAlbum(c *gin.Context) {
	var input schemas.CreateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := services.AlbumService{}
	albumID, err := s.CreateAlbum(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": albumID, "created": true})
}

func GetAllAlbums(c *gin.Context) {
	s := services.AlbumService{}
	albums, err := s.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func GetAlbumById(c *gin.Context) {
	s := services.AlbumService{}
	album, err := s.GetByID(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func UpdateAlbum(c *gin.Context) {
	var input schemas.UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := services.AlbumService{}
	album, err := s.Update(c.Param("id"), input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func DeleteAlbum(c *gin.Context) {
	s := services.AlbumService{}
	err := s.DeleteByID(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"deleted": true})
}
