package services

import (
	"example/web-service-gin/pkg/models"
	"example/web-service-gin/pkg/repositories"
	"example/web-service-gin/pkg/schemas"
)

type AlbumService struct{}

func (s AlbumService) CreateAlbum(input schemas.CreateAlbumInput) (int64, error) {
	repo := repositories.AlbumRepository{}
	return repo.CreateAlbum(input)
}

func (s AlbumService) GetAll() ([]models.Album, error) {
	repo := repositories.AlbumRepository{}
	return repo.GetAll()
}

func (s AlbumService) GetByID(id string) (models.Album, error) {
	repo := repositories.AlbumRepository{}
	return repo.GetByID(id)
}

func (s AlbumService) Update(id string, input schemas.UpdateAlbumInput) (models.Album, error) {
	repo := repositories.AlbumRepository{}
	return repo.Update(id, input)
}

func (s AlbumService) DeleteByID(id string) error {
	repo := repositories.AlbumRepository{}
	return repo.DeleteByID(id)
}
