package middleware_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"

	helperfakes "github.com/BrandonWade/enako/api/helpers/fakes"
)

var _ = Describe("AuthenticateMiddleware", func() {
	var (
		logger    *logrus.Logger
		store     *helperfakes.FakeCookieStorer
		session   *helperfakes.FakeSessionStorer
		stack     *middleware.MiddlewareStack
		mw        middleware.Middleware
		decorator func(http.ResponseWriter, *http.Request)
		w         *httptest.ResponseRecorder
		r         *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		store = &helperfakes.FakeCookieStorer{}
		session = &helperfakes.FakeSessionStorer{}
		stack = middleware.NewMiddlewareStack(logger, store)

		decorator = func(w http.ResponseWriter, r *http.Request) {}
		mw = stack.Authenticate()
		w = httptest.NewRecorder()
	})

	Describe("Authenticate", func() {
		Context("when authenticating an incoming request", func() {
			It("returns an error if an error is encountered retrieving the session", func() {
				store.IsAuthenticatedReturns(false, errors.New("session error"))
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				resBody := fmt.Sprintf(`[{"text":"%s","type":"error"}]`, helpers.ErrorFetchingSession())

				handler := mw(decorator)
				handler(w, r)

				Expect(store.IsAuthenticatedCallCount()).To(Equal(1))
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if the session is not authenticated", func() {
				store.IsAuthenticatedReturns(false, nil)
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				resBody := fmt.Sprintf(`[{"text":"%s","type":"error"}]`, helpers.ErrorUserNotAuthenticated())

				handler := mw(decorator)
				handler(w, r)

				Expect(store.IsAuthenticatedCallCount()).To(Equal(1))
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("calls the next function with no error if the session is authenticated", func() {
				accountID := int64(1)
				session.GetReturns(1)
				store.GetReturns(session, nil)
				store.IsAuthenticatedReturns(true, nil)
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				ctx := context.WithValue(r.Context(), middleware.ContextAccountIDKey, accountID)
				r = r.WithContext(ctx)

				handler := mw(decorator)
				handler(w, r)

				Expect(store.IsAuthenticatedCallCount()).To(Equal(1))
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEmpty())
			})
		})
	})
})
