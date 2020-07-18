package middleware_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/sirupsen/logrus"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
)

var _ = Describe("DecodeExpenseMiddleware", func() {
	var (
		logger    *logrus.Logger
		store     helpers.CookieStorer
		stack     *middleware.Stack
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
		mw = stack.DecodeExpense()
		w = httptest.NewRecorder()
	})

	Describe("DecodeExpense", func() {
		Context("when decoding an Expense from an incoming request", func() {
			It("returns an error if a malformed payload is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/expenses", strings.NewReader("{foo}"))
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidExpensePayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error", func() {
				categoryID := int64(5)
				description := "test"
				amount := float64(1234)
				expenseDate := "2020-06-01"

				payload := models.Expense{CategoryID: categoryID, Description: description, Amount: amount, ExpenseDate: expenseDate}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})

			It("stores the Expense in the request Context", func() {
				categoryID := int64(5)
				description := "test"
				amount := float64(1234)
				expenseDate := "2020-06-01"

				payload := models.Expense{CategoryID: categoryID, Description: description, Amount: amount, ExpenseDate: expenseDate}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/expenses", bytes.NewBuffer(payloadJSON))
				decorator = func(w http.ResponseWriter, r *http.Request) {
					expense, ok := r.Context().Value(middleware.ContextExpenseKey).(models.Expense)

					Expect(expense).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
						"CategoryID":  Equal(categoryID),
						"Description": Equal(description),
						"Amount":      Equal(amount),
						"ExpenseDate": Equal(expenseDate),
					}))
					Expect(ok).To(BeTrue())
				}

				handler := mw(decorator)
				handler(w, r)
			})
		})
	})
})
