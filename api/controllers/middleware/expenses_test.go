package middleware_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/controllers/middleware"
	"github.com/BrandonWade/enako/api/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("ExpenseMiddleware", func() {
	var (
		logger    *logrus.Logger
		decorator func(http.ResponseWriter, *models.Expense)
		w         *httptest.ResponseRecorder
		r         *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		decorator = func(w http.ResponseWriter, expense *models.Expense) {}

		w = httptest.NewRecorder()
	})

	Describe("DecodeExpense", func() {

		Context("when decoding an expense", func() {

			It("returns an error if a malformed payload is submitted", func() {
				handler := middleware.DecodeExpense(logger, decorator)
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrInvalidExpensePayload)
				r = httptest.NewRequest("POST", "/v1/expenses", strings.NewReader("{foo}"))

				handler(w, r)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns nothing if there were no errors", func() {
				handler := middleware.DecodeExpense(logger, decorator)
				r = httptest.NewRequest("POST", "/v1/expenses", strings.NewReader(`{"ID": 1}`))

				handler(w, r)
				Expect(w.Body.String()).To(BeEmpty())
			})
		})
	})
})
