package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services/fakes"
	"github.com/gorilla/mux"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExpenseController", func() {
	var (
		expenseService    *fakes.FakeExpenseService
		expenseController controllers.ExpenseController
		expenses          []models.UserExpense
		w                 *httptest.ResponseRecorder
		r                 *http.Request
	)

	BeforeEach(func() {
		expenseService = &fakes.FakeExpenseService{}
		expenseController = controllers.NewExpenseController(expenseService)

		expenses = []models.UserExpense{
			models.UserExpense{ID: 1, UserAccountID: 100, ExpenseType: "type 1", ExpenseTypeID: 333, ExpenseCategory: "category 1", ExpenseCategoryID: 4444, ExpenseDescription: "test description", ExpenseAmount: 100, ExpenseDate: "2019-01-01", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.UserExpense{ID: 2, UserAccountID: 200, ExpenseType: "type 2", ExpenseTypeID: 444, ExpenseCategory: "category 2", ExpenseCategoryID: 5555, ExpenseDescription: "test description", ExpenseAmount: 200, ExpenseDate: "2019-01-01", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
			models.UserExpense{ID: 3, UserAccountID: 300, ExpenseType: "type 3", ExpenseTypeID: 555, ExpenseCategory: "category 3", ExpenseCategoryID: 6666, ExpenseDescription: "test description", ExpenseAmount: 300, ExpenseDate: "2019-01-01", CreatedAt: "2019-01-01 00:00:00", UpdatedAt: "2019-01-01 00:00:00"},
		}

		w = httptest.NewRecorder()
	})

	Describe("GetExpenses", func() {

		Context("when requesting the list of expenses", func() {

			It("returns an error when the expense service returns an error", func() {
				expenseService.GetExpensesReturns([]models.UserExpense{}, errors.New("service error"))
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

			It("returns an error if a malformed payload is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", strings.NewReader("{foo}"))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrInvalidExpensePayload)

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid expense type id is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 0, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))
				resBody := `{"errors":["ExpenseTypeID: less than min"]}`

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid expense category id is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 0, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))
				resBody := `{"errors":["ExpenseCategoryID: less than min"]}`

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid expense amount is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 0, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))
				resBody := `{"errors":["ExpenseAmount: less than min"]}`

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid date is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "0000-00-00"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))
				resBody := `{"errors":["ExpenseDate: invalid date"]}`

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if a malformed date is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))
				resBody := `{"errors":["ExpenseDate: invalid date"]}`

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if one was encountered while communicating with the expense service", func() {
				expenseService.CreateExpenseReturns(0, errors.New("service error"))
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrCreatingExpense)

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns the info for the created expense with no error", func() {
				expenseID := int64(100)

				expenseService.CreateExpenseReturns(expenseID, nil)
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))

				response := models.UserExpense{ID: expenseID, ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				responseJSON, _ := json.Marshal(response)

				expenseController.CreateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(w.Body.String()).To(BeEquivalentTo(string(responseJSON) + "\n"))
			})
		})
	})

	Describe("UpdateExpense", func() {

		Context("when updating an expense", func() {

			It("returns an error if an invalid expense id is provided", func() {
				r = httptest.NewRequest("PUT", "/v1/expenses/id", nil)
				r = mux.SetURLVars(r, map[string]string{"id": "foo"})
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrInvalidExpenseID)

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if a malformed payload is submitted", func() {
				r = httptest.NewRequest("PUT", "/v1/expenses/id", strings.NewReader("{foo}"))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrInvalidExpensePayload)

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid expense type id is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 0, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := `{"errors":["ExpenseTypeID: less than min"]}`

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid expense category id is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 0, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := `{"errors":["ExpenseCategoryID: less than min"]}`

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid expense amount is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 0, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := `{"errors":["ExpenseAmount: less than min"]}`

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid date is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "0000-00-00"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := `{"errors":["ExpenseDate: invalid date"]}`

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if a malformed date is submitted", func() {
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := `{"errors":["ExpenseDate: invalid date"]}`

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if one is encountered while communicating with the expense service", func() {
				expenseService.UpdateExpenseReturns(0, errors.New("service error"))
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrUpdatingExpense)

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if the given expense could not be found", func() {
				expenseService.UpdateExpenseReturns(0, nil)
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrNoExpensesUpdated)

				expenseController.UpdateExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns the info for the updated expense with no error", func() {
				expenseID := int64(123)
				typeID := int64(5)
				categoryID := int64(5)
				description := "test"
				amount := int64(1234)
				date := "2019-01-01"

				expenseService.UpdateExpenseReturns(1, nil)
				payload := models.UserExpense{ExpenseTypeID: typeID, ExpenseCategoryID: categoryID, ExpenseDescription: description, ExpenseAmount: amount, ExpenseDate: date}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": strconv.FormatInt(expenseID, 10)})

				response := models.UserExpense{ID: expenseID, ExpenseTypeID: typeID, ExpenseCategoryID: categoryID, ExpenseDescription: description, ExpenseAmount: amount, ExpenseDate: date}
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
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
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
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
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
				payload := models.UserExpense{ExpenseTypeID: 5, ExpenseCategoryID: 5, ExpenseDescription: "test", ExpenseAmount: 1234, ExpenseDate: "2019-01-01"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("DELETE", "/v1/expenses/id", bytes.NewBuffer(payloadJSON))
				r = mux.SetURLVars(r, map[string]string{"id": "123"})

				expenseController.DeleteExpense(w, r)
				Expect(w.Code).To(Equal(http.StatusNoContent))
			})
		})
	})
})
