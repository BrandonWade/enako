package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services/fakes"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	helperfakes "github.com/BrandonWade/enako/api/helpers/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExpenseController", func() {
	var (
		logger            *logrus.Logger
		store             *helperfakes.FakeCookieStorer
		expenseService    *fakes.FakeExpenseService
		expenseController controllers.ExpenseController
		expenses          []models.Expense
		w                 *httptest.ResponseRecorder
		r                 *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		store = &helperfakes.FakeCookieStorer{}

		expenseService = &fakes.FakeExpenseService{}
		expenseController = controllers.NewExpenseController(logger, store, expenseService)

		expenses = []models.Expense{
			models.Expense{ID: 1, UserAccountID: 100, Category: "category 1", CategoryID: 4444, Description: "test description", Amount: 100, ExpenseDate: "2019-01-01", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.Expense{ID: 2, UserAccountID: 200, Category: "category 2", CategoryID: 5555, Description: "test description", Amount: 200, ExpenseDate: "2019-01-01", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.Expense{ID: 3, UserAccountID: 300, Category: "category 3", CategoryID: 6666, Description: "test description", Amount: 300, ExpenseDate: "2019-01-01", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
		}

		w = httptest.NewRecorder()
	})

	Describe("GetExpenses", func() {
		Context("when requesting the list of expenses", func() {
			It("returns an error when the expense service returns an error", func() {
				expenseService.GetExpensesReturns([]models.Expense{}, errors.New("service error"))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrFetchingExpenses)

				expenseController.GetExpenses(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(string(resBody) + "\n"))
			})

			It("returns the list of expenses with no error", func() {
				expenseService.GetExpensesReturns(expenses, nil)
				resBody, _ := json.Marshal(expenses)

				expenseController.GetExpenses(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(BeEquivalentTo(string(resBody) + "\n"))
			})
		})
	})

	Describe("CreateExpense", func() {
		Context("when creating a new expense", func() {
			// TODO: Migrate these tests cases to be middleware tests instead
			// It("returns an error if an invalid expense category id is submitted", func() {
			// 	r = httptest.NewRequest("POST", "/v1/expenses", nil)
			// 	payload := models.Expense{CategoryID: 0, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
			// 	resBody := `{"errors":["CategoryID: less than min"]}`
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)

			// 	expenseController.CreateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			// TODO: Add test case for verifying Description isn't empty

			// It("returns an error if an invalid expense amount is submitted", func() {
			// 	r = httptest.NewRequest("POST", "/v1/expenses", nil)
			// 	payload := models.Expense{CategoryID: 5, Description: "test", Amount: 0, ExpenseDate: "2019-01-01"}
			// 	resBody := `{"errors":["Amount: less than min"]}`
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)

			// 	expenseController.CreateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			// It("returns an error if an invalid date is submitted", func() {
			// 	r = httptest.NewRequest("POST", "/v1/expenses", nil)
			// 	payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "0000-00-00"}
			// 	resBody := `{"errors":["ExpenseDate: invalid date"]}`
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)

			// 	expenseController.CreateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			// It("returns an error if a malformed date is submitted", func() {
			// 	r = httptest.NewRequest("POST", "/v1/expenses", nil)
			// 	payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01"}
			// 	resBody := `{"errors":["ExpenseDate: invalid date"]}`
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)

			// 	expenseController.CreateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			It("returns an error if one was encountered while communicating with the expense service", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", nil)
				expenseService.CreateExpenseReturns(0, errors.New("service error"))
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrCreatingExpense)
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns the info for the created expense with no error", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", nil)
				expenseID := int64(100)

				expenseService.CreateExpenseReturns(expenseID, nil)
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)

				response := models.Expense{ID: expenseID, CategoryID: 5, Description: "test", Amount: 12.34, ExpenseDate: "2019-01-01"}
				responseJSON, _ := json.Marshal(response)

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(w.Body.String()).To(BeEquivalentTo(string(responseJSON) + "\n"))
			})
		})
	})

	Describe("UpdateExpense", func() {
		Context("when updating an expense", func() {
			// TODO: Migrate these tests cases to be middleware tests instead
			// It("returns an error if an invalid expense id is provided", func() {
			// 	payload := models.Expense{CategoryID: 1, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
			// 	r = httptest.NewRequest("PUT", "/v1/expenses/id", nil)
			// 	r = mux.SetURLVars(r, map[string]string{"id": "foo"})
			// 	resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrInvalidExpenseID)
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)

			// 	expenseController.UpdateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusBadRequest))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			// It("returns an error if an invalid expense category id is submitted", func() {
			// 	payload := models.Expense{CategoryID: 0, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
			// 	payloadJSON, _ := json.Marshal(payload)

			// 	r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
			// 	r = mux.SetURLVars(r, map[string]string{"id": "123"})
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)
			// 	resBody := `{"errors":["CategoryID: less than min"]}`

			// 	expenseController.UpdateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			// TODO: Add test case for verifying Description isn't empty

			// It("returns an error if an invalid expense amount is submitted", func() {
			// 	payload := models.Expense{CategoryID: 5, Description: "test", Amount: 0, ExpenseDate: "2019-01-01"}
			// 	payloadJSON, _ := json.Marshal(payload)

			// 	r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
			// 	r = mux.SetURLVars(r, map[string]string{"id": "123"})
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)
			// 	resBody := `{"errors":["Amount: less than min"]}`

			// 	expenseController.UpdateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			// It("returns an error if an invalid date is submitted", func() {
			// 	payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "0000-00-00"}
			// 	payloadJSON, _ := json.Marshal(payload)

			// 	r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
			// 	r = mux.SetURLVars(r, map[string]string{"id": "123"})
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)
			// 	resBody := `{"errors":["ExpenseDate: invalid date"]}`

			// 	expenseController.UpdateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			// It("returns an error if a malformed date is submitted", func() {
			// 	payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01"}
			// 	payloadJSON, _ := json.Marshal(payload)

			// 	r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
			// 	r = mux.SetURLVars(r, map[string]string{"id": "123"})
			// 	ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
			// 	r = r.WithContext(ctx)
			// 	resBody := `{"errors":["ExpenseDate: invalid date"]}`

			// 	expenseController.UpdateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			// TODO: This test case likely needs additional work
			// It("returns an error if a malformed payload is submitted", func() {
			// 	r = httptest.NewRequest("PUT", "/v1/expenses/id", strings.NewReader("{foo}"))
			// 	r = mux.SetURLVars(r, map[string]string{"id": "123"})
			// 	resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrInvalidExpensePayload)

			// 	expenseController.UpdateExpense(w, r)
			// 	Expect(w.Code).To(Equal(http.StatusBadRequest))
			// 	Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			// })

			It("returns an error if one is encountered while communicating with the expense service", func() {
				expenseService.UpdateExpenseReturns(0, errors.New("service error"))
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrUpdatingExpense)

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if the given expense could not be found", func() {
				expenseService.UpdateExpenseReturns(0, nil)
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrNoExpensesUpdated)

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns the info for the updated expense with no error", func() {
				expenseID := int64(123)
				categoryID := int64(5)
				description := "test"
				amount := float64(1234)
				date := "2019-01-01"

				expenseService.UpdateExpenseReturns(1, nil)
				payload := models.Expense{CategoryID: categoryID, Description: description, Amount: amount, ExpenseDate: date}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": strconv.FormatInt(expenseID, 10)})
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)

				response := models.Expense{ID: expenseID, CategoryID: categoryID, Description: description, Amount: amount / 100, ExpenseDate: date}
				responseJSON, _ := json.Marshal(response)

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(BeEquivalentTo(string(responseJSON) + "\n"))
			})
		})
	})

	Describe("DeleteExpense", func() {
		Context("when deleting an expense", func() {
			It("returns an error if an invalid expense id is provided", func() {
				r = httptest.NewRequest("DELETE", "/v1/expenses/id", nil)
				r = mux.SetURLVars(r, map[string]string{"id": "foo"})
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrInvalidExpenseID)

				expenseController.DeleteExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if one is encountered while communicating with the expense service", func() {
				expenseService.DeleteExpenseReturns(0, errors.New("service error"))
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("DELETE", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrDeletingExpense)

				expenseController.DeleteExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if the given expense could not be found", func() {
				expenseService.DeleteExpenseReturns(0, nil)
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("DELETE", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrNoExpensesDeleted)

				expenseController.DeleteExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns content and no error if the expense was deleted", func() {
				expenseService.DeleteExpenseReturns(1, nil)
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("DELETE", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})

				expenseController.DeleteExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusNoContent))
			})
		})
	})
})
