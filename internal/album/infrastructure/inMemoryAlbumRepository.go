package infrastructure

import (
	"errors"

	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
)

type InMemoryAlbumRepository struct {
	db *[]domain.Album
}

func NewInMemoryAlbumRepository(db *[]domain.Album) domain.AlbumRepository {
	return &InMemoryAlbumRepository{db: db}
}

func (r *InMemoryAlbumRepository) GetAlbums() (*[]domain.Album, error) {
	if len(*r.db) == 0 {
		return nil, errors.New("there are no albums in database")
	}

	return r.db, nil
}
