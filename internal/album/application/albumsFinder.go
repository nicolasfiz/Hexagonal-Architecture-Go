package application

import (
	"github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
)

type AlbumFinder struct {
	repository domain.AlbumRepository
}

func NewAlbumFinder(repo domain.AlbumRepository) *AlbumFinder {
	return &AlbumFinder{repository: repo}
}

func (f *AlbumFinder) Run() (*[]domain.Album, error) {
	return f.repository.GetAlbums()
}
