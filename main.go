package main

import (
	"example/web-service-gin/pkg/config"
	"example/web-service-gin/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", controllers.FindAlbums)
	router.GET("/albums/:id", controllers.FindAlbumById)
	router.POST("/albums", controllers.CreateAlbum)
	router.PATCH("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	return router
}

func main() {
	router := SetUpRouter()
	config.ConnectDatabase()
	router.Run("localhost:8000")
}
