package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/services"
)

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
		// TODO: Handle
	}

	json.NewEncoder(w).Encode(types)
}
