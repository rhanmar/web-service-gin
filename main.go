package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/albums", controllers.FindAlbums)
	router.POST("/albums", controllers.CreateAlbum)

	router.Run("localhost:8000")
}
