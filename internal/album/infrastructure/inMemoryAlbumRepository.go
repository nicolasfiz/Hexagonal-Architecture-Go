package infrastructure

import (
	"errors"

	albumDomain "github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
)

type InMemoryAlbumRepository struct {
	db *[]albumDomain.Album
}

func (r *InMemoryAlbumRepository) GetAlbums() (*[]albumDomain.Album, error) {
	if len(*r.db) == 0 {
		return nil, errors.New("there are no albums in database")
	}

	return r.db, nil
}
