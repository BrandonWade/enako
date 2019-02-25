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

var ErrFetchingTypes = errors.New("error fetching types")

//go:generate counterfeiter -o fakes/fake_type_controller.go . TypeController
type TypeController interface {
	GetTypes(w http.ResponseWriter, r *http.Request)
}

type typeController struct {
	logger  *logrus.Logger
	store   helpers.CookieStorer
	service services.TypeService
}

func NewTypeController(logger *logrus.Logger, store helpers.CookieStorer, service services.TypeService) TypeController {
	return &typeController{
		logger,
		store,
		service,
	}
}

func (t *typeController) GetTypes(w http.ResponseWriter, r *http.Request) {
	types, err := t.service.GetTypes()
	if err != nil {
		t.logger.WithFields(logrus.Fields{
			"method": "TypeController.GetTypes",
			"err":    err.Error(),
		}).Error(ErrFetchingTypes)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(ErrFetchingTypes))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types)
	return
}
