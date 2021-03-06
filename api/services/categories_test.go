package services_test

import (
	"errors"
	"io/ioutil"

	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories/fakes"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CategoryService", func() {
	var (
		logger          *logrus.Logger
		categoryRepo    *fakes.FakeCategoryRepository
		categoryService services.CategoryService

		categoryList = []models.Category{
			models.Category{
				ID:        111,
				Name:      "test category",
				CreatedAt: "2018-01-01 00:00:00",
				UpdatedAt: "2018-01-01 00:00:00",
			},
			models.Category{
				ID:        222,
				Name:      "another test category",
				CreatedAt: "2018-01-01 00:00:00",
				UpdatedAt: "2018-01-01 00:00:00",
			},
			models.Category{
				ID:        333,
				Name:      "yet another test category",
				CreatedAt: "2018-01-01 00:00:00",
				UpdatedAt: "2018-01-01 00:00:00",
			},
		}
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		categoryRepo = &fakes.FakeCategoryRepository{}
		categoryService = services.NewCategoryService(logger, categoryRepo)
	})

	Describe("GetCategories", func() {
		Context("when requesting the list of categories", func() {
			It("returns an error if an error is encountered", func() {
				categoryRepo.GetCategoriesReturns([]models.Category{}, errors.New("repo error"))

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
