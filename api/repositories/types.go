package repositories

import (
	"github.com/BrandonWade/enako/api/models"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type TypesRepository interface {
	GetTypes() ([]models.Type, error)
}

type typesRepository struct {
	DB *sqlx.DB
}

func NewTypesRepository(DB *sqlx.DB) TypesRepository {
	return &typesRepository{
		DB,
	}
}

func (t *typesRepository) GetTypes() ([]models.Type, error) {
	types := []models.Type{}

	err := t.DB.Select(&types, `SELECT *
        FROM types;
    `)
	if err != nil {
		return []models.Type{}, err
	}

	return types, nil
}
