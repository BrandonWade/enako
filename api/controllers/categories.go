package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/models"
)

type CategoriesController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}

type categoriesController struct {
	categories []models.Category
}

func NewCategoriesController() CategoriesController {
	return &categoriesController{
		[]models.Category{ // TODO: Hardcoded for testing
			models.Category{
				ID:   1,
				Name: "Food",
			},
			models.Category{
				ID:   2,
				Name: "Entertainment",
			},
			models.Category{
				ID:   3,
				Name: "Transportation",
			},
			models.Category{
				ID:   4,
				Name: "Clothing",
			},
			models.Category{
				ID:   5,
				Name: "Technology",
			},
			models.Category{
				ID:   6,
				Name: "Health",
			},
		},
	}
}

func (c *categoriesController) GetCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.categories)
}
