package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/infrastructure"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/pkg/db"
)

func InitAlbumRoutes(r *gin.Engine, db *db.Database) {
	albumGroup := r.Group("/album")

	albumHandler := infrastructure.NewAlbumHandler(db)

	albumGroup.GET("/", albumHandler.GetAlbums)
}
