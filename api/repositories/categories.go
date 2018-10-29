package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type CategoriesRepository interface {
	GetCategories() ([]models.Category, error)
}

type categoriesRepository struct {
	DB *sqlx.DB
}

func NewCategoriesRepository(DB *sqlx.DB) CategoriesRepository {
	return &categoriesRepository{
		DB,
	}
}

func (c *categoriesRepository) GetCategories() ([]models.Category, error) {
	categories := []models.Category{}

	err := c.DB.Select(&categories, `SELECT *
        FROM categories;
    `)
	if err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}
