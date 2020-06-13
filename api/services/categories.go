package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/sirupsen/logrus"
)

// CategoryService an interface for working with categories.
//go:generate counterfeiter -o fakes/fake_category_service.go . CategoryService
type CategoryService interface {
	GetCategories() ([]models.Category, error)
}

type categoryService struct {
	logger *logrus.Logger
	repo   repositories.CategoryRepository
}

// NewCategoryService returns a new instance of a CategoryService.
func NewCategoryService(logger *logrus.Logger, repo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		logger,
		repo,
	}
}

// GetCategories retrieves the list of categories.
func (c *categoryService) GetCategories() ([]models.Category, error) {
	return c.repo.GetCategories()
}
