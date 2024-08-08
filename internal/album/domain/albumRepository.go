package domain

type AlbumRepository interface {
	GetAlbums() (*[]Album, error)
}
