package main

import (
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

	albumHandler := album.AlbumHandler{Db: &db.Db}

	albumGroup.GET("/", albumHandler.GetAlbums)

	r.Run(":8080")
}
