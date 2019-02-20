package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/sirupsen/logrus"
)

//go:generate counterfeiter -o fakes/fake_category_service.go . CategoryService
type CategoryService interface {
	GetCategories() ([]models.ExpenseCategory, error)
}

type categoryService struct {
	logger *logrus.Logger
	repo   repositories.CategoryRepository
}

func NewCategoryService(logger *logrus.Logger, repo repositories.CategoryRepository) CategoryService {
	return &categoryService{
		logger,
		repo,
	}
}

func (c *categoryService) GetCategories() ([]models.ExpenseCategory, error) {
	return c.repo.GetCategories()
}
