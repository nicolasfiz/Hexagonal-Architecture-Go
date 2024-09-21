package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/pkg/db"
)

func InitRoutes(r *gin.Engine, db *db.Database) {
	InitAlbumRoutes(r, db)
	InitHealthRoutes(r)
}
