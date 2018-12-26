package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

//go:generate counterfeiter -o fakes/fake_category_service.go . CategoryService
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
