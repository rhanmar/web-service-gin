package main

import (
	"example/web-service-gin/pkg/config"
	"example/web-service-gin/pkg/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", controllers.GetAllAlbums)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.POST("/albums", controllers.CreateAlbum)
	router.PATCH("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	return router
}

func main() {
	router := SetUpRouter()
	db, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
	}
	router.Run("localhost:8000")

	defer db.Close()
}
