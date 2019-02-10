package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services"

	log "github.com/sirupsen/logrus"
)

var errFetchingTypes = errors.New("error fetching types")

//go:generate counterfeiter -o fakes/fake_type_controller.go . TypeController
type TypeController interface {
	GetTypes(w http.ResponseWriter, r *http.Request)
}

type typeController struct {
	service services.TypeService
}

func NewTypeController(service services.TypeService) TypeController {
	return &typeController{
		service,
	}
}

func (t *typeController) GetTypes(w http.ResponseWriter, r *http.Request) {
	types, err := t.service.GetTypes()
	if err != nil {
		log.WithFields(log.Fields{
			"method": "TypeController.GetTypes",
			"err":    err.Error(),
		}).Error(errFetchingTypes)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.NewAPIError(errFetchingTypes))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(types)
	return
}
