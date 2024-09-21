package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/application"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/pkg/db"
)

type AlbumHandler struct {
	db         *db.Database
	repository domain.AlbumRepository
}

func NewAlbumHandler(db *db.Database) *AlbumHandler {
	repository := NewInMemoryAlbumRepository(db)
	return &AlbumHandler{db, repository}
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	finder := application.NewAlbumFinder(h.repository)
	albums, err := finder.Run()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, *albums)
}
