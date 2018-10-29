package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

type TypesService interface {
	GetTypes() ([]models.Type, error)
}

type typesService struct {
	repo repositories.TypesRepository
}

func NewTypesService(repo repositories.TypesRepository) TypesService {
	return &typesService{
		repo,
	}
}

func (t *typesService) GetTypes() ([]models.Type, error) {
	return t.repo.GetTypes()
}
