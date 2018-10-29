package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type CategoryRepository interface {
	GetCategories() ([]models.Category, error)
}

type categoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(DB *sqlx.DB) CategoryRepository {
	return &categoryRepository{
		DB,
	}
}

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
