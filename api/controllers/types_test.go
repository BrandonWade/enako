package controllers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services/fakes"
	"github.com/sirupsen/logrus"

	helperfakes "github.com/BrandonWade/enako/api/helpers/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TypeController", func() {

	var (
		logger         *logrus.Logger
		store          *helperfakes.FakeCookieStorer
		typeService    *fakes.FakeTypeService
		typeController controllers.TypeController
		types          []models.ExpenseType
		w              *httptest.ResponseRecorder
		r              *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		store = &helperfakes.FakeCookieStorer{}

		typeService = &fakes.FakeTypeService{}
		typeController = controllers.NewTypeController(logger, store, typeService)

		types = []models.ExpenseType{
			models.ExpenseType{ID: 1, TypeName: "type 1", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.ExpenseType{ID: 2, TypeName: "type 2", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.ExpenseType{ID: 3, TypeName: "type 3", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
		}

		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/v1/types", nil)
	})

	Describe("GetTypes", func() {

		Context("when requesting the list of types", func() {

			It("returns an error if an error is encountered", func() {
				typeService.GetTypesReturns([]models.ExpenseType{}, errors.New("service error"))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrFetchingTypes)

				typeController.GetTypes(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns the list of types with no error", func() {
				typeService.GetTypesReturns(types, nil)
				resBody, _ := json.Marshal(types)

				typeController.GetTypes(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(BeEquivalentTo(string(resBody) + "\n"))
			})
		})
	})
})
