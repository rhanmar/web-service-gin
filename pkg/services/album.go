package services

import (
	"example/web-service-gin/pkg/config"
	"example/web-service-gin/pkg/models"
	"example/web-service-gin/pkg/schemas"
)

type AlbumService struct{}

func (s AlbumService) CreateAlbum(input schemas.CreateAlbumInput) (models.Album, error) {
	album := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}
	config.DB.Create(&album)
	return album, nil
}

func (s AlbumService) GetAll() ([]models.Album, error) {
	var albums []models.Album
	config.DB.Find(&albums)
	return albums, nil
}

func (s AlbumService) GetByID(id string) (models.Album, error) {
	var album models.Album
	if err := config.DB.Where("id = ?", id).First(&album).Error; err != nil {
		return album, err
	}
	return album, nil
}

func (s AlbumService) Update(id string, input schemas.UpdateAlbumInput) (models.Album, error) {
	album, err := s.GetByID(id)
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

func (s AlbumService) DeleteByID(id string) error {
	album, err := s.GetByID(id)
	if err != nil {
		return err
	}

	config.DB.Delete(&album)
	return nil
}
