package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", createAlbum)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.PATCH("/albums/:id", updateAlbumByID)

	router.Run("localhost:8000")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func createAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if id == a.ID {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"info": "Album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if id == a.ID {
			albums = deleteAlbum(i, albums)
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"info": "Album not found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}
	var albumIndex int
	for i, a := range albums {
		if id == a.ID {
			albumIndex = i
			break
		}
	}
	albums[albumIndex].Title = updatedAlbum.Title
	albums[albumIndex].Price = updatedAlbum.Price
	c.IndentedJSON(http.StatusOK, albums[albumIndex])
}
