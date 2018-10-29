package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

type CategoriesService interface {
	GetCategories() ([]models.Category, error)
}

type categoriesService struct {
	repo repositories.CategoriesRepository
}

func NewCategoriesService(repo repositories.CategoriesRepository) CategoriesService {
	return &categoriesService{
		repo,
	}
}

func (c *categoriesService) GetCategories() ([]models.Category, error) {
	return c.repo.GetCategories()
}
