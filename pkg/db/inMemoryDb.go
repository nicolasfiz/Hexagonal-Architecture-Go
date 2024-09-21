package db

import albumDomain "github.com/nicolasfiz/Hexagonal-Architecture-Go/internal/album/domain"

type Database struct {
	AlbumData *[]albumDomain.Album
}

var database = []albumDomain.Album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetDatabase() *Database {
	return &Database{AlbumData: &database}
}
