package services

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories"
)

//go:generate counterfeiter -o fakes/fake_type_service.go . TypeService
type TypeService interface {
	GetTypes() ([]models.ExpenseType, error)
}

type typeService struct {
	repo repositories.TypeRepository
}

func NewTypeService(repo repositories.TypeRepository) TypeService {
	return &typeService{
		repo,
	}
}

func (t *typeService) GetTypes() ([]models.ExpenseType, error) {
	return t.repo.GetTypes()
}
