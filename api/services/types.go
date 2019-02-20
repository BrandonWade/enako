package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
	"github.com/sirupsen/logrus"
)

//go:generate counterfeiter -o fakes/fake_type_service.go . TypeService
type TypeService interface {
	GetTypes() ([]models.ExpenseType, error)
}

type typeService struct {
	logger *logrus.Logger
	repo   repositories.TypeRepository
}

func NewTypeService(logger *logrus.Logger, repo repositories.TypeRepository) TypeService {
	return &typeService{
		logger,
		repo,
	}
}

func (t *typeService) GetTypes() ([]models.ExpenseType, error) {
	return t.repo.GetTypes()
}
