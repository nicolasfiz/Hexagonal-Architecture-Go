package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	albumApplication "github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/application"
	albumDomain "github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
)

type AlbumHandler struct {
	Db *[]albumDomain.Album
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	repo := &InMemoryAlbumRepository{db: h.Db}
	finder := &albumApplication.AlbumFinder{Repository: repo}
	albums, err := finder.Run()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, *albums)
}
