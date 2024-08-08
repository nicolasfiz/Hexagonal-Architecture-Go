package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	album "github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/infrastructure"
	db "github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/db"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Albums
	albumGroup := r.Group("/album")

	albumHandler := album.NewAlbumHandler(db.GetDatabase())

	albumGroup.GET("/", albumHandler.GetAlbums)

	fmt.Println("🚀 Server running at http://localhost:8080 🚀")
	r.Run(":8080")
}
