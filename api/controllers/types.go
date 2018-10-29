package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/services"
)

type TypesController interface {
	GetTypes(w http.ResponseWriter, r *http.Request)
}

type typesController struct {
	service services.TypesService
}

func NewTypesController(service services.TypesService) TypesController {
	return &typesController{
		service,
	}
}

func (t *typesController) GetTypes(w http.ResponseWriter, r *http.Request) {
	types, err := t.service.GetTypes()
	if err != nil {
		// TODO: Handle
	}

	json.NewEncoder(w).Encode(types)
}
