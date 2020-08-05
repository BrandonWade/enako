package middleware_test

import (
	"bytes"
	"encoding/json"
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
	"github.com/onsi/gomega/gstruct"
	"github.com/sirupsen/logrus"
)

var _ = Describe("DecodeChangePasswordMiddleware", func() {
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
		mw = stack.DecodeChangePassword()
		w = httptest.NewRecorder()
	})

	Describe("DecodeChangePassword", func() {
		Context("when decoding a ChangePassword from an incoming request", func() {
			It("returns an error if a malformed payload is submitted", func() {
				r = httptest.NewRequest("PUT", "/v1/accounts/password/change", strings.NewReader("{foo}"))
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidChangePasswordPayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error", func() {
				currentPassword := "testpassword123"
				newPassword := "testpassword12345"
				confirmPassword := "testpassword12345"

				payload := models.ChangePassword{CurrentPassword: currentPassword, NewPassword: newPassword, ConfirmPassword: confirmPassword}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/accounts/password/change", bytes.NewBuffer(payloadJSON))

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})

			It("stores the decoded ChangePassword payload in the request context", func() {
				currentPassword := "testpassword123"
				newPassword := "testpassword12345"
				confirmPassword := "testpassword12345"

				payload := models.ChangePassword{CurrentPassword: currentPassword, NewPassword: newPassword, ConfirmPassword: confirmPassword}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("PUT", "/v1/accounts/password/change", bytes.NewBuffer(payloadJSON))
				decorator = func(w http.ResponseWriter, r *http.Request) {
					changePassword, ok := r.Context().Value(middleware.ContextChangePasswordKey).(models.ChangePassword)

					Expect(changePassword).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
						"CurrentPassword": Equal(currentPassword),
						"NewPassword":     Equal(newPassword),
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
