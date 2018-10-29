package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/services"
)

type CategoriesController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}

type categoriesController struct {
	service services.CategoriesService
}

func NewCategoriesController(service services.CategoriesService) CategoriesController {
	return &categoriesController{
		service,
	}
}

func (c *categoriesController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.service.GetCategories()
	if err != nil {
		// TODO: Handle
	}

	json.NewEncoder(w).Encode(categories)
}
