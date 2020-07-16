package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"
)

// CategoryController a controller for working with categories.
//go:generate counterfeiter -o fakes/fake_category_controller.go . CategoryController
type CategoryController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}

type categoryController struct {
	logger  *logrus.Logger
	store   helpers.CookieStorer
	service services.CategoryService
}

// NewCategoryController returns a new instance of a CategoryController.
func NewCategoryController(logger *logrus.Logger, store helpers.CookieStorer, service services.CategoryService) CategoryController {
	return &categoryController{
		logger,
		store,
		service,
	}
}

// GetCategories returns the list of categories.
func (c *categoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.service.GetCategories()
	if err != nil {
		c.logger.WithField("method", "CategoryController.GetCategories").Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.MessagesFromErrors(helpers.ErrorFetchingCategories()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
	return
}
