package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/application"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
)

type AlbumHandler struct {
	db *[]domain.Album
}

func NewAlbumHandler(db *[]domain.Album) *AlbumHandler {
	return &AlbumHandler{db}
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	repo := NewInMemoryAlbumRepository(h.db)
	finder := application.NewAlbumFinder(repo)
	albums, err := finder.Run()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, *albums)
}
