package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

type CategoryService interface {
	GetCategories() ([]models.Category, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repo,
	}
}

func (c *categoryService) GetCategories() ([]models.Category, error) {
	return c.repo.GetCategories()
}
