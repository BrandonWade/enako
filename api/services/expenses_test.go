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

var _ = Describe("ExpenseService", func() {
	var (
		logger         *logrus.Logger
		expenseRepo    *fakes.FakeExpenseRepository
		expenseService services.ExpenseService

		accountID   = int64(123456)
		expenseList = []models.Expense{
			models.Expense{
				ID:          1,
				AccountID:   12345,
				Category:    "test category",
				CategoryID:  123,
				Description: "test description",
				Amount:      111,
				ExpenseDate: "2018-01-01 00:00:00",
				CreatedAt:   "2018-01-01 00:00:00",
				UpdatedAt:   "2018-01-01 00:00:00",
			},
			models.Expense{
				ID:          2,
				AccountID:   1328904,
				Category:    "another test category",
				CategoryID:  2340985,
				Description: "another test description",
				Amount:      222,
				ExpenseDate: "2018-01-01 00:00:00",
				CreatedAt:   "2018-01-01 00:00:00",
				UpdatedAt:   "2018-01-01 00:00:00",
			},
			models.Expense{
				ID:          3,
				AccountID:   17486329,
				Category:    "yet another test category",
				CategoryID:  123678,
				Description: "yet another test description",
				Amount:      333,
				ExpenseDate: "2018-01-01 00:00:00",
				CreatedAt:   "2018-01-01 00:00:00",
				UpdatedAt:   "2018-01-01 00:00:00",
			},
		}
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		expenseRepo = &fakes.FakeExpenseRepository{}
		expenseService = services.NewExpenseService(logger, expenseRepo)
	})

	Describe("GetExpenses", func() {
		Context("when requesting the list of expenses", func() {
			It("returns an error if an error is encountered", func() {
				expenseRepo.GetExpensesReturns([]models.Expense{}, errors.New("repo error"))

				expenses, err := expenseService.GetExpenses(accountID)
				Expect(expenseRepo.GetExpensesCallCount()).To(Equal(1))
				Expect(expenses).To(BeEmpty())
				Expect(err).To(HaveOccurred())
			})

			It("returns a list of expenses with no error", func() {
				expenseRepo.GetExpensesReturns(expenseList, nil)

				expenses, err := expenseService.GetExpenses(accountID)
				Expect(expenseRepo.GetExpensesCallCount()).To(Equal(1))
				Expect(expenses).To(Equal(expenseList))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("CreateExpense", func() {
		Context("when creating a new expense", func() {
			var (
				expenseID = int64(1928736)
				expense   = &models.Expense{
					ID:          123312,
					Description: "test expense",
				}
			)

			It("returns an error if an error is encountered", func() {
				expenseRepo.CreateExpenseReturns(0, errors.New("repo error"))

				ID, err := expenseService.CreateExpense(accountID, expense)
				Expect(expenseRepo.CreateExpenseCallCount()).To(Equal(1))
				Expect(ID).To(Equal(int64(0)))
				Expect(err).To(HaveOccurred())
			})

			It("returns the id of the new expense row with no error", func() {
				expenseRepo.CreateExpenseReturns(expenseID, nil)

				ID, err := expenseService.CreateExpense(accountID, expense)
				Expect(expenseRepo.CreateExpenseCallCount()).To(Equal(1))
				Expect(ID).To(Equal(expenseID))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("UpdateExpense", func() {
		Context("when updating an existing expense with the given id", func() {
			var (
				expenseID = int64(235476)
				expense   = &models.Expense{
					ID:          123312,
					Description: "test expense",
				}
			)

			It("returns an error if an error is encountered", func() {
				expenseRepo.UpdateExpenseReturns(0, errors.New("repo error"))

				count, err := expenseService.UpdateExpense(expenseID, accountID, expense)
				Expect(expenseRepo.UpdateExpenseCallCount()).To(Equal(1))
				Expect(count).To(Equal(int64(0)))
				Expect(err).To(HaveOccurred())
			})

			It("returns the number of the updated rows with no error", func() {
				expenseRepo.UpdateExpenseReturns(1, nil)

				count, err := expenseService.UpdateExpense(expenseID, accountID, expense)
				Expect(expenseRepo.UpdateExpenseCallCount()).To(Equal(1))
				Expect(count).To(Equal(int64(1)))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("DeleteExpense", func() {
		Context("when deleting an existing expense with the given id", func() {
			var (
				expenseID = int64(637485)
			)

			It("returns an error if an error is encountered", func() {
				expenseRepo.DeleteExpenseReturns(0, errors.New("repo error"))

				count, err := expenseService.DeleteExpense(expenseID, accountID)
				Expect(expenseRepo.DeleteExpenseCallCount()).To(Equal(1))
				Expect(count).To(Equal(int64(0)))
				Expect(err).To(HaveOccurred())
			})

			It("returns the number of the deleted rows with no error", func() {
				expenseRepo.DeleteExpenseReturns(1, nil)

				count, err := expenseService.DeleteExpense(expenseID, accountID)
				Expect(expenseRepo.DeleteExpenseCallCount()).To(Equal(1))
				Expect(count).To(Equal(int64(1)))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
