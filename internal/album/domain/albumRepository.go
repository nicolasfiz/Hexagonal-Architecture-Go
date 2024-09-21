package domain

type AlbumRepository interface {
	GetAlbums() (*[]Album, error)
	CreateAlbum(album *Album) (*Album, error)
}
