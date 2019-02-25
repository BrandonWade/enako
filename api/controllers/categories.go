package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"
)

var ErrFetchingCategories = errors.New("error fetching categories")

//go:generate counterfeiter -o fakes/fake_category_controller.go . CategoryController
type CategoryController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
}

type categoryController struct {
	logger  *logrus.Logger
	store   helpers.CookieStorer
	service services.CategoryService
}

func NewCategoryController(logger *logrus.Logger, store helpers.CookieStorer, service services.CategoryService) CategoryController {
	return &categoryController{
		logger,
		store,
		service,
	}
}

func (c *categoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.service.GetCategories()
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"method": "CategoryController.GetCategories",
			"err":    err.Error(),
		}).Error(ErrFetchingCategories)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrFetchingCategories))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
	return
}
