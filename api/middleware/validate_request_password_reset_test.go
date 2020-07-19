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

var _ = Describe("ValidateRequestPasswordResetMiddleware", func() {
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
		mw = stack.ValidateRequestPasswordReset()
		w = httptest.NewRecorder()
	})

	Describe("ValidateRequestPasswordReset", func() {
		Context("when validating a RequestPasswordReset from an incoming request", func() {
			It("returns an error if one is encountered retrieving the RequestPasswordReset from the request Context", func() {
				r = httptest.NewRequest("POST", "/v1/password", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidRequestPasswordResetPayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if an invalid email is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/password", nil)
				payload := models.RequestPasswordReset{Email: "invalid@@invalid.com"}
				resBody := `{"messages":[{"text":"Email: invalid email","type":"error"}]}`
				ctx := context.WithValue(r.Context(), middleware.ContextRequestPasswordResetKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error if the request is valid", func() {
				r = httptest.NewRequest("POST", "/v1/password", nil)
				payload := models.RequestPasswordReset{Email: "email@test.com"}
				ctx := context.WithValue(r.Context(), middleware.ContextRequestPasswordResetKey, payload)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})
		})
	})
})
