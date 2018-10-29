package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type TypeRepository interface {
	GetTypes() ([]models.Type, error)
}

type typeRepository struct {
	DB *sqlx.DB
}

func NewTypeRepository(DB *sqlx.DB) TypeRepository {
	return &typeRepository{
		DB,
	}
}

func (t *typeRepository) GetTypes() ([]models.Type, error) {
	types := []models.Type{}

	err := t.DB.Select(&types, `SELECT *
        FROM types;
    `)
	if err != nil {
		return []models.Type{}, err
	}

	return types, nil
}
