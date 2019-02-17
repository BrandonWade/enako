package services_test

import (
	"errors"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories/fakes"
	"github.com/BrandonWade/enako/api/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CategoryService", func() {

	var (
		categoryRepo    *fakes.FakeCategoryRepository
		categoryService services.CategoryService

		categoryList = []models.ExpenseCategory{
			models.ExpenseCategory{
				ID:           111,
				CategoryName: "test category",
				CreatedAt:    "2018-01-01 00:00:00",
				UpdatedAt:    "2018-01-01 00:00:00",
			},
			models.ExpenseCategory{
				ID:           222,
				CategoryName: "another test category",
				CreatedAt:    "2018-01-01 00:00:00",
				UpdatedAt:    "2018-01-01 00:00:00",
			},
			models.ExpenseCategory{
				ID:           333,
				CategoryName: "yet another test category",
				CreatedAt:    "2018-01-01 00:00:00",
				UpdatedAt:    "2018-01-01 00:00:00",
			},
		}
	)

	BeforeEach(func() {
		categoryRepo = &fakes.FakeCategoryRepository{}
		categoryService = services.NewCategoryService(categoryRepo)
	})

	Describe("GetCategories", func() {

		Context("when requesting the list of categories", func() {

			It("returns an error if an error is encountered", func() {
				categoryRepo.GetCategoriesReturns([]models.ExpenseCategory{}, errors.New("repo error"))

				categories, err := categoryService.GetCategories()
				Expect(categoryRepo.GetCategoriesCallCount()).To(Equal(1))
				Expect(categories).To(BeEmpty())
				Expect(err).To(HaveOccurred())
			})

			It("returns the list of categories with no error", func() {
				categoryRepo.GetCategoriesReturns(categoryList, nil)

				categories, err := categoryService.GetCategories()
				Expect(categoryRepo.GetCategoriesCallCount()).To(Equal(1))
				Expect(categories).To(Equal(categoryList))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
