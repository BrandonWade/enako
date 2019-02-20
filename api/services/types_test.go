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

var _ = Describe("TypeService", func() {

	var (
		logger      *logrus.Logger
		typeRepo    *fakes.FakeTypeRepository
		typeService services.TypeService

		typeList = []models.ExpenseType{
			models.ExpenseType{
				ID:        111,
				TypeName:  "test type",
				CreatedAt: "2018-01-01 00:00:00",
				UpdatedAt: "2018-01-01 00:00:00",
			},
			models.ExpenseType{
				ID:        222,
				TypeName:  "another test type",
				CreatedAt: "2018-01-01 00:00:00",
				UpdatedAt: "2018-01-01 00:00:00",
			},
			models.ExpenseType{
				ID:        333,
				TypeName:  "yet another test type",
				CreatedAt: "2018-01-01 00:00:00",
				UpdatedAt: "2018-01-01 00:00:00",
			},
		}
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		typeRepo = &fakes.FakeTypeRepository{}
		typeService = services.NewTypeService(logger, typeRepo)
	})

	Describe("GetTypes", func() {

		Context("when requesting the list of types", func() {

			It("returns an error if an error is encountered", func() {
				typeRepo.GetTypesReturns([]models.ExpenseType{}, errors.New("repo error"))

				types, err := typeService.GetTypes()
				Expect(typeRepo.GetTypesCallCount()).To(Equal(1))
				Expect(types).To(BeEmpty())
				Expect(err).To(HaveOccurred())
			})

			It("returns the list of types with no error", func() {
				typeRepo.GetTypesReturns(typeList, nil)

				types, err := typeService.GetTypes()
				Expect(typeRepo.GetTypesCallCount()).To(Equal(1))
				Expect(types).To(Equal(typeList))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
