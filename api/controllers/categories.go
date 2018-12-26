package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/services"
)

//go:generate counterfeiter -o fakes/fake_category_controller.go . CategoryController
type CategoryController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}

type categoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &categoryController{
		service,
	}
}

func (c *categoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.service.GetCategories()
	if err != nil {
		// TODO: Handle
	}

	json.NewEncoder(w).Encode(categories)
}
