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

var _ = Describe("ValidatePasswordResetMiddleware", func() {
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
		mw = stack.ValidatePasswordReset()
		w = httptest.NewRecorder()
	})

	Describe("ValidatePasswordReset", func() {
		Context("when validating a PasswordReset from an incoming request", func() {
			It("returns an error if one is encountered retrieving the PasswordReset from the request Context", func() {
				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidPasswordResetPayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that is too short is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				payload := models.PasswordReset{Password: "password123", ConfirmPassword: "testpassword123"}
				resBody := `{"messages":[{"text":"Password: password must be minimum 15 characters","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextPasswordResetKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that is too long is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				payload := models.PasswordReset{Password: "thisisareallylongpasswordthatistoolongandwillfailvalidation", ConfirmPassword: "thisisareallylongpassword"}
				resBody := `{"messages":[{"text":"Password: password must be maximum 50 characters","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextPasswordResetKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that contains invalid characters is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				payload := models.PasswordReset{Password: "_-123testpassword456-_", ConfirmPassword: "123testpassword456"}
				resBody := `{"messages":[{"text":"Password: password may only contain alphanumeric characters and the following symbols: _ ! @ # $ % ^ *","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextPasswordResetKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if the passwords do not match", func() {
				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				payload := models.PasswordReset{Password: "testpassword123", ConfirmPassword: "testpassword1234"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorPasswordsDoNotMatch())
				ctx := context.WithValue(r.Context(), middleware.ContextPasswordResetKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error if the request is valid", func() {
				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				payload := models.PasswordReset{Password: "testpassword123", ConfirmPassword: "testpassword123"}
				ctx := context.WithValue(r.Context(), middleware.ContextPasswordResetKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})
		})
	})
})
