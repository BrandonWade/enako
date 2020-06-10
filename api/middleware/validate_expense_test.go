package middleware_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/validation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("ValidateExpenseMiddleware", func() {
	var (
		logger    *logrus.Logger
		store     helpers.CookieStorer
		stack     *middleware.MiddlewareStack
		mw        middleware.Middleware
		decorator func(http.ResponseWriter, *http.Request)
		w         *httptest.ResponseRecorder
		r         *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		store = &helpers.CookieStore{}
		stack = middleware.NewMiddlewareStack(logger, store)

		decorator = func(w http.ResponseWriter, r *http.Request) {}
		mw = stack.ValidateExpense()
		w = httptest.NewRecorder()

		validation.InitValidator()
	})

	Describe("ValidateExpense", func() {
		Context("when validating an Expense from an incoming request", func() {
			It("returns an error if an error is encountered retrieving the Expense from the request Context", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", nil)
				resBody := `{"errors":["invalid expense payload"]}`

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if an invalid expense category id is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", nil)
				payload := models.Expense{CategoryID: 0, Description: "test", Amount: 1234, ExpenseDate: "2019-01-01"}
				resBody := `{"errors":["CategoryID: less than min"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if an empty description is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", nil)
				payload := models.Expense{CategoryID: 5, Description: "", Amount: 1234, ExpenseDate: "2019-01-01"}
				resBody := `{"errors":["Description: zero value"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if an invalid expense amount is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", nil)
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 0, ExpenseDate: "2019-01-01"}
				resBody := `{"errors":["Amount: less than min"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if an invalid date is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", nil)
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "0000-00-00"}
				resBody := `{"errors":["ExpenseDate: invalid date"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a malformed date is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", nil)
				payload := models.Expense{CategoryID: 5, Description: "test", Amount: 1234, ExpenseDate: "2019-01"}
				resBody := `{"errors":["ExpenseDate: invalid date"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextExpenseKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})
		})
	})
})
