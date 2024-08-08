package application

import (
	AlbumDomain "github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"
)

type AlbumFinder struct {
	Repository AlbumDomain.AlbumRepository
}

func (f *AlbumFinder) Run() (*[]AlbumDomain.Album, error) {
	return f.Repository.GetAlbums()
}
