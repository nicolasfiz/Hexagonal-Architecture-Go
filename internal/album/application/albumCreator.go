package application

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
)

type AlbumCreator struct {
	repository domain.AlbumRepository
}

func NewAlbumCreator(repo domain.AlbumRepository) *AlbumCreator {
	return &AlbumCreator{repository: repo}
}

func (c *AlbumCreator) Run(album *domain.Album) (*domain.Album, error) {

	if album.Title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}

	if album.Price == 0 {
		return nil, fmt.Errorf("price cannot be 0 or empty")
	}

	album.Id = uuid.New().String()

	return c.repository.CreateAlbum(album)
}
