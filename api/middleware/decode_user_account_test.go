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

var _ = Describe("DecodeCreateAccountMiddleware", func() {
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
		mw = stack.DecodeCreateAccount()
		w = httptest.NewRecorder()
	})

	Describe("DecodeCreateAccount", func() {
		Context("when decoding a CreateAccount from an incoming request", func() {
			It("returns an error if a malformed payload is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", strings.NewReader("{foo}"))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, helpers.ErrorInvalidAccountPayload())

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error", func() {
				username := "test"
				email := "test@example.com"
				password := "testpassword123"
				confirmPassword := "testpassword123"

				payload := models.CreateAccount{Username: username, Email: email, Password: password, ConfirmPassword: confirmPassword}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))

				handler := mw(decorator)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})

			It("stores the decoded CreateAccount payload in the request context", func() {
				username := "test"
				email := "test@example.com"
				password := "testpassword123"
				confirmPassword := "testpassword123"

				payload := models.CreateAccount{Username: username, Email: email, Password: password, ConfirmPassword: confirmPassword}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))
				decorator = func(w http.ResponseWriter, r *http.Request) {
					createAccount, ok := r.Context().Value(middleware.ContextCreateAccountKey).(models.CreateAccount)

					Expect(createAccount).To(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
						"Username":        Equal(username),
						"Email":           Equal(email),
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
