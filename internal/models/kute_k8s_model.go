// internal/models/your_model.go
package models

import "errors"

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Dimpu", Artist: "Jassi", Price: 56.99},
	{ID: "2", Title: "Chikku", Artist: "Yogansh", Price: 17.99},
}

// Error constants
var (
	ErrAlbumNotFound = errors.New("album not found")
)

// GetAlbums returns the list of all albums.
func GetAlbums() []Album {
	return albums
}

// AddAlbum adds a new album to the collection.
func AddAlbum(album Album) {
	albums = append(albums, album)
}

// FindAlbumByID searches for an album by ID.
func FindAlbumByID(id string) (Album, error) {
	for _, album := range albums {
		if album.ID == id {
			return album, nil
		}
	}
	return Album{}, ErrAlbumNotFound
}
