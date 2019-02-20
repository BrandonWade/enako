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

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CategoryController", func() {

	var (
		logger             *logrus.Logger
		categoryService    *fakes.FakeCategoryService
		categoryController controllers.CategoryController
		categories         []models.ExpenseCategory
		w                  *httptest.ResponseRecorder
		r                  *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		categoryService = &fakes.FakeCategoryService{}
		categoryController = controllers.NewCategoryController(logger, categoryService)

		categories = []models.ExpenseCategory{
			models.ExpenseCategory{ID: 1, CategoryName: "category 1", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.ExpenseCategory{ID: 2, CategoryName: "category 2", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.ExpenseCategory{ID: 3, CategoryName: "category 3", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
		}

		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/v1/categories", nil)
	})

	Describe("GetCategories", func() {

		Context("when requesting the list of categories", func() {

			It("returns an error if an error is encountered", func() {
				categoryService.GetCategoriesReturns([]models.ExpenseCategory{}, errors.New("service error"))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrFetchingCategories)

				categoryController.GetCategories(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns the list of categories with no error", func() {
				categoryService.GetCategoriesReturns(categories, nil)
				resBody, _ := json.Marshal(categories)

				categoryController.GetCategories(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(BeEquivalentTo(string(resBody) + "\n"))
			})
		})
	})
})
