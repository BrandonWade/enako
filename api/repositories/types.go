package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate counterfeiter -o fakes/fake_type_repository.go . TypeRepository
type TypeRepository interface {
	GetTypes() ([]models.ExpenseType, error)
}

type typeRepository struct {
	DB *sqlx.DB
}

func NewTypeRepository(DB *sqlx.DB) TypeRepository {
	return &typeRepository{
		DB,
	}
}

func (t *typeRepository) GetTypes() ([]models.ExpenseType, error) {
	types := []models.ExpenseType{}

	err := t.DB.Select(&types, `SELECT *
        FROM expense_types;
    `)
	if err != nil {
		return []models.ExpenseType{}, err
	}

	return types, nil
}
