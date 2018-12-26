package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate counterfeiter -o fakes/fake_category_repository.go . CategoryRepository
type CategoryRepository interface {
	GetCategories() ([]models.ExpenseCategory, error)
}

type categoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(DB *sqlx.DB) CategoryRepository {
	return &categoryRepository{
		DB,
	}
}

func (c *categoryRepository) GetCategories() ([]models.ExpenseCategory, error) {
	categories := []models.ExpenseCategory{}

	err := c.DB.Select(&categories, `SELECT *
        FROM expense_categories;
    `)
	if err != nil {
		return []models.ExpenseCategory{}, err
	}

	return categories, nil
}
