package repositories

import (
	"example/web-service-gin/pkg/config"
	"example/web-service-gin/pkg/models"
	"example/web-service-gin/pkg/schemas"
)

type AlbumRepository struct{}

func (repo AlbumRepository) CreateAlbum(input schemas.CreateAlbumInput) (models.Album, error) {
	album := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}
	config.DB.Create(&album)
	return album, nil
}

func (repo AlbumRepository) GetAll() ([]models.Album, error) {
	var albums []models.Album
	config.DB.Find(&albums)
	return albums, nil
}

func (repo AlbumRepository) GetByID(id string) (models.Album, error) {
	var album models.Album
	if err := config.DB.Where("id = ?", id).First(&album).Error; err != nil {
		return album, err
	}
	return album, nil
}

func (repo AlbumRepository) Update(id string, input schemas.UpdateAlbumInput) (models.Album, error) {
	album, err := repo.GetByID(id)
	if err != nil {
		return album, err
	}

	updatedAlbum := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}

	config.DB.Model(&album).Updates(&updatedAlbum)
	return album, nil
}

func (repo AlbumRepository) DeleteByID(id string) error {
	album, err := repo.GetByID(id)
	if err != nil {
		return err
	}

	config.DB.Delete(&album)
	return nil
}
