package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"
)

//go:generate counterfeiter -o fakes/fake_category_controller.go . CategoryController
type CategoryController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}

type categoryController struct {
	logger  *logrus.Logger
	store   helpers.CookieStorer
	service services.CategoryService
}

// NewCategoryController ...
func NewCategoryController(logger *logrus.Logger, store helpers.CookieStorer, service services.CategoryService) CategoryController {
	return &categoryController{
		logger,
		store,
		service,
	}
}

// GetCategories ...
func (c *categoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.service.GetCategories()
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"method": "CategoryController.GetCategories",
			"err":    err.Error(),
		}).Error(helpers.ErrorFetchingCategories())

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(helpers.ErrorFetchingCategories()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
	return
}
