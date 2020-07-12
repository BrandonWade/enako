package middleware_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("ValidateCreateAccountMiddleware", func() {
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
		mw = stack.ValidateCreateAccount()
		w = httptest.NewRecorder()
	})

	Describe("ValidateCreateAccount", func() {
		Context("when validating a CreateAccount from an incoming request", func() {
			It("returns an error if an error is encountered retrieving the CreateAccount from the request Context", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, helpers.ErrorInvalidAccountPayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if an invalid email is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				payload := models.CreateAccount{Email: "invalid@@invalid.com", Password: "testpassword123", ConfirmPassword: "testpassword123"}
				resBody := `{"errors":["Email: invalid email"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextCreateAccountKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that is too short is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				payload := models.CreateAccount{Email: "test@email.com", Password: "password", ConfirmPassword: "password1234567890"}
				resBody := `{"errors":["Password: password must be minimum 15 characters"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextCreateAccountKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that is too long is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				payload := models.CreateAccount{Email: "test@email.com", Password: "thisisareallylongpasswordthatistoolongandwillfailvalidation", ConfirmPassword: "thisisareallylongpassword"}
				resBody := `{"errors":["Password: password must be maximum 50 characters"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextCreateAccountKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that contains invalid characters is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				payload := models.CreateAccount{Email: "test@email.com", Password: "_-123testpassword456-_", ConfirmPassword: "123testpassword456"}
				resBody := `{"errors":["Password: password may only contain alphanumeric characters and the following symbols: _ ! @ # $ % ^ *"]}`
				ctx := context.WithValue(r.Context(), middleware.ContextCreateAccountKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if the passwords do not match", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				payload := models.CreateAccount{Email: "email@test.com", Password: "testpassword123", ConfirmPassword: "testpassword1234"}
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, helpers.ErrorPasswordsDoNotMatch())
				ctx := context.WithValue(r.Context(), middleware.ContextCreateAccountKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})
		})
	})
})
