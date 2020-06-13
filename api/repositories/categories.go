package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

// CategoryRepository an interface for working with categories.
//go:generate counterfeiter -o fakes/fake_category_repository.go . CategoryRepository
type CategoryRepository interface {
	GetCategories() ([]models.Category, error)
}

type categoryRepository struct {
	DB *sqlx.DB
}

// NewCategoryRepository returns a new instance of a CategoryRepository.
func NewCategoryRepository(DB *sqlx.DB) CategoryRepository {
	return &categoryRepository{
		DB,
	}
}

// GetCategories retrieves the list of categories.
func (c *categoryRepository) GetCategories() ([]models.Category, error) {
	categories := []models.Category{}

	err := c.DB.Select(&categories, `SELECT *
        FROM categories;
    `)
	if err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}
