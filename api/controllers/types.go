package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BrandonWade/enako/api/models"
)

type TypesController interface {
	GetTypes(w http.ResponseWriter, r *http.Request)
}

type typesController struct {
	types []models.Type
}

func NewTypesController() TypesController {
	return &typesController{
		[]models.Type{ // TODO: Hardcoded for testing
			models.Type{
				ID:   1,
				Name: "General",
			},
			models.Type{
				ID:   2,
				Name: "Unnecessary",
			},
			models.Type{
				ID:   3,
				Name: "Recurring",
			},
		},
	}
}

func (t *typesController) GetTypes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(t.types)
}
