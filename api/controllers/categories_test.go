package controllers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services/fakes"
	"github.com/sirupsen/logrus"

	helperfakes "github.com/BrandonWade/enako/api/helpers/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CategoryController", func() {
	var (
		logger             *logrus.Logger
		store              *helperfakes.FakeCookieStorer
		categoryService    *fakes.FakeCategoryService
		categoryController controllers.CategoryController
		categories         []models.Category
		w                  *httptest.ResponseRecorder
		r                  *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		store = &helperfakes.FakeCookieStorer{}

		categoryService = &fakes.FakeCategoryService{}
		categoryController = controllers.NewCategoryController(logger, store, categoryService)

		categories = []models.Category{
			models.Category{ID: 1, Name: "category 1", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.Category{ID: 2, Name: "category 2", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.Category{ID: 3, Name: "category 3", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
		}

		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/v1/categories", nil)
	})

	Describe("GetCategories", func() {
		Context("when requesting the list of categories", func() {
			It("returns an error if an error is encountered", func() {
				categoryService.GetCategoriesReturns([]models.Category{}, errors.New("service error"))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, helpers.ErrorFetchingCategories())

				categoryController.GetCategories(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns the list of categories with no error", func() {
				categoryService.GetCategoriesReturns(categories, nil)
				resBody, _ := json.Marshal(categories)

				categoryController.GetCategories(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})
		})
	})
})
