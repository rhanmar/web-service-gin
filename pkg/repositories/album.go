package repositories

import (
	"example/web-service-gin/pkg/config"
	"example/web-service-gin/pkg/models"
	"example/web-service-gin/pkg/schemas"
)

type AlbumRepository struct{}

func (repo AlbumRepository) CreateAlbum(input schemas.CreateAlbumInput) (int64, error) {
	result, err := config.DB.Exec(
		"INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id",
		input.Title,
		input.Artist,
		input.Price,
	)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func (repo AlbumRepository) GetAll() ([]models.Album, error) {
	var albums []models.Album
	err := config.DB.Select(&albums, "SELECT * FROM albums")
	return albums, err
}

func (repo AlbumRepository) GetByID(id string) (models.Album, error) {
	var album models.Album
	err := config.DB.Get(&album, "SELECT * FROM albums WHERE id =  $1", id)
	return album, err
}

func (repo AlbumRepository) Update(id string, input schemas.UpdateAlbumInput) (models.Album, error) {
	album, err := repo.GetByID(id)
	if err != nil {
		return album, err
	}
	_, err = config.DB.Exec(
		"UPDATE albums SET title=$1, artist=$2, price=$3 WHERE id = $4",
		input.Title,
		input.Artist,
		input.Price,
		id,
	)
	return album, err
}

func (repo AlbumRepository) DeleteByID(id string) error {
	_, err := repo.GetByID(id)
	if err != nil {
		return err
	}
	_, err = config.DB.Exec("DELETE FROM albums WHERE id = $1", id)
	return err
}
