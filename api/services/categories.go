package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

type CategoryService interface {
	GetCategories() ([]models.ExpenseCategory, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		repo,
	}
}

func (c *categoryService) GetCategories() ([]models.ExpenseCategory, error) {
	return c.repo.GetCategories()
}
