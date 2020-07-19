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

var _ = Describe("DecodeRequestPasswordResetMiddleware", func() {
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
		mw = stack.DecodeRequestPasswordReset()
		w = httptest.NewRecorder()
	})

	Describe("DecodeRequestPasswordReset", func() {
		Context("when decoding a RequestPasswordReset from an incoming request", func() {
			It("returns an error if a malformed payload is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/password", strings.NewReader("{foo}"))
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidRequestPasswordResetPayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error", func() {
				email := "foo@bar.net"

				payload := models.RequestPasswordReset{Email: email}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/password", bytes.NewBuffer(payloadJSON))

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})

			It("stores the RequestPasswordReset in the request Context", func() {
				email := "foo@bar.net"

				payload := models.RequestPasswordReset{Email: email}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/password", bytes.NewBuffer(payloadJSON))
				decorator = func(w http.ResponseWriter, r *http.Request) {
					requestPasswordReset, ok := r.Context().Value(middleware.ContextRequestPasswordResetKey).(models.RequestPasswordReset)

					Expect(requestPasswordReset).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
						"Email": Equal(email),
					}))
					Expect(ok).To(BeTrue())
				}

				handler := mw(decorator)
				handler(w, r)
			})
		})
	})
})
