package infrastructure

import (
	"errors"
	"fmt"

	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/pkg/db"
)

type InMemoryAlbumRepository struct {
	db *db.Database
}

func NewInMemoryAlbumRepository(db *db.Database) domain.AlbumRepository {
	return &InMemoryAlbumRepository{db: db}
}

func (r *InMemoryAlbumRepository) GetAlbums() (*[]domain.Album, error) {
	if len(*r.db.AlbumData) == 0 {
		return nil, errors.New("there are no albums in database")
	}

	return r.db.AlbumData, nil
}

func (r *InMemoryAlbumRepository) CreateAlbum(album *domain.Album) (*domain.Album, error) {
	if album == nil {
		return nil, fmt.Errorf("you cannot create an empty album")
	}

	*r.db.AlbumData = append(*r.db.AlbumData, *album)

	return album, nil
}
