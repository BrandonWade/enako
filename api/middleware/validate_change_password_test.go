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

var _ = Describe("ValidateChangePasswordMiddleware", func() {
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
		mw = stack.ValidateChangePassword()
		w = httptest.NewRecorder()
	})

	Describe("ValidateChangePassword", func() {
		Context("when validating a ChangePassword from an incoming request", func() {
			It("returns an error if one is encountered retrieving the ChangePassword from the request context", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidChangePasswordPayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that is too short is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "shortpassword", ConfirmPassword: "testpassword12345"}
				resBody := `{"messages":[{"text":"NewPassword: password must be minimum 15 characters","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that is too long is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "thisisareallylongpasswordthatistoolongandwillfailvalidation", ConfirmPassword: "testpassword12345"}
				resBody := `{"messages":[{"text":"NewPassword: password must be maximum 50 characters","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a password that contains invalid characters is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "_-testpassword12345-_", ConfirmPassword: "testpassword12345"}
				resBody := `{"messages":[{"text":"NewPassword: password may only contain alphanumeric characters and the following symbols: _ ! @ # $ % ^ *","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a confirm password that is too short is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "testpassword12345", ConfirmPassword: "shortpassword"}
				resBody := `{"messages":[{"text":"ConfirmPassword: password must be minimum 15 characters","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a confirm password that is too long is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "testpassword12345", ConfirmPassword: "thisisareallylongpasswordthatistoolongandwillfailvalidation"}
				resBody := `{"messages":[{"text":"ConfirmPassword: password must be maximum 50 characters","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if a confirm password that contains invalid characters is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "testpassword12345", ConfirmPassword: "_-testpassword12345-_"}
				resBody := `{"messages":[{"text":"ConfirmPassword: password may only contain alphanumeric characters and the following symbols: _ ! @ # $ % ^ *","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if the current password is the same as the new password", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "testpassword123", ConfirmPassword: "testpassword123"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorPasswordsShouldNotMatch())
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if the new and confirm passwords do not match", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "testpassword12345", ConfirmPassword: "testpassword123456"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorPasswordsDoNotMatch())
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error if the request is valid", func() {
				r = httptest.NewRequest("POST", "/v1/accounts/password/change", nil)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "testpassword12345", ConfirmPassword: "testpassword12345"}
				ctx := context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})
		})
	})
})
