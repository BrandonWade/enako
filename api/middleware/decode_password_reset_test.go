package middleware_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/models"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/sirupsen/logrus"
)

var _ = Describe("DecodePasswordResetMiddleware", func() {
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
		mw = stack.DecodePasswordReset()
		w = httptest.NewRecorder()
	})

	Describe("DecodePasswordReset", func() {
		Context("when decoding a PasswordReset from an incoming request", func() {
			It("returns an error if a malformed payload is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/password/reset", strings.NewReader("{foo}"))
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidPasswordResetPayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error", func() {
				password := "testpassword123"
				confirmPassword := "testpassword123"

				payload := models.PasswordReset{Password: password, ConfirmPassword: confirmPassword}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/password/reset", bytes.NewBuffer(payloadJSON))

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})

			It("stores the PasswordReset in the request Context", func() {
				password := "testpassword123"
				confirmPassword := "testpassword123"

				payload := models.PasswordReset{Password: password, ConfirmPassword: confirmPassword}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/password/reset", bytes.NewBuffer(payloadJSON))
				decorator = func(w http.ResponseWriter, r *http.Request) {
					passwordReset, ok := r.Context().Value(middleware.ContextPasswordResetKey).(models.PasswordReset)

					Expect(passwordReset).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
						"Password":        Equal(password),
						"ConfirmPassword": Equal(confirmPassword),
					}))
					Expect(ok).To(BeTrue())
				}

				handler := mw(decorator)
				handler(w, r)
			})
		})
	})
})
