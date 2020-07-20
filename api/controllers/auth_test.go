package controllers_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/helpers"
	helperfakes "github.com/BrandonWade/enako/api/helpers/fakes"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("AuthController", func() {
	var (
		logger         *logrus.Logger
		store          *helperfakes.FakeCookieStorer
		session        *helperfakes.FakeSessionStorer
		accountService *fakes.FakeAccountService
		authController controllers.AuthController
		w              *httptest.ResponseRecorder
		r              *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		store = &helperfakes.FakeCookieStorer{}
		session = &helperfakes.FakeSessionStorer{}

		accountService = &fakes.FakeAccountService{}
		authController = controllers.NewAuthController(logger, store, accountService)

		w = httptest.NewRecorder()
	})

	Describe("CSRF", func() {
		Context("when attempting to generate a new CSRF token", func() {
			It("sets a response header containing the CSRF token", func() {
				r = httptest.NewRequest("HEAD", "/v1/csrf", nil)

				authController.CSRF(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.HeaderMap).To(HaveKey("X-Csrf-Token"))
			})
		})
	})

	Describe("Login", func() {
		Context("when attempting to login", func() {
			It("returns an error when one occurred while retrieving the session", func() {
				store.GetReturns(nil, errors.New("session error"))
				r = httptest.NewRequest("POST", "/v1/login", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorFetchingSession())

				authController.Login(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns an error if one occurred while retrieving the Account with the login credentials from the request context", func() {
				store.GetReturns(session, nil)
				r = httptest.NewRequest("POST", "/v1/login", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidAccountPayload())

				authController.Login(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns an error if one is returned from the account service because the account has not been activated and an email was resent", func() {
				accountEmail := "foo@bar.net"

				store.GetReturns(session, nil)
				r = httptest.NewRequest("POST", "/v1/login", nil)
				accountService.VerifyAccountReturns(0, helpers.ErrorActivationEmailResent())
				payload := models.Account{Email: accountEmail, Password: "testpassword123"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"info"}]}`, helpers.MessageActivationEmailSent(accountEmail))
				ctx := context.WithValue(r.Context(), middleware.ContextLoginKey, payload)
				r = r.WithContext(ctx)

				authController.Login(w, r)
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns an error if one is returned from the account service because the account has not been activated", func() {
				accountEmail := "foo@bar.net"

				store.GetReturns(session, nil)
				r = httptest.NewRequest("POST", "/v1/login", nil)
				accountService.VerifyAccountReturns(0, helpers.ErrorAccountNotActivated())
				payload := models.Account{Email: accountEmail, Password: "testpassword123"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorAccountNotActivated())
				ctx := context.WithValue(r.Context(), middleware.ContextLoginKey, payload)
				r = r.WithContext(ctx)

				authController.Login(w, r)
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns an error if an invalid username or password is provided", func() {
				accountEmail := "foo@bar.net"

				store.GetReturns(session, nil)
				r = httptest.NewRequest("POST", "/v1/login", nil)
				accountService.VerifyAccountReturns(0, helpers.ErrorInvalidEmailOrPassword())
				payload := models.Account{Email: accountEmail, Password: "testpassword123"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorInvalidEmailOrPassword())
				ctx := context.WithValue(r.Context(), middleware.ContextLoginKey, payload)
				r = r.WithContext(ctx)

				authController.Login(w, r)
				Expect(w.Code).To(Equal(http.StatusUnauthorized))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("marks the session as authenticated and returns without an error", func() {
				accountID := int64(1)
				accountEmail := "foo@bar.net"

				store.GetReturns(session, nil)
				r = httptest.NewRequest("POST", "/v1/login", nil)
				accountService.VerifyAccountReturns(accountID, nil)
				payload := models.Account{Email: accountEmail, Password: "testpassword123"}
				resBody := fmt.Sprintf(`{"id":%d,"email":"%s"}`, accountID, accountEmail)
				ctx := context.WithValue(r.Context(), middleware.ContextLoginKey, payload)
				r = r.WithContext(ctx)

				authController.Login(w, r)
				Expect(session.SetCallCount()).To(Equal(2))
				setArgsFirst1, setArgsSecond1 := session.SetArgsForCall(0)
				setArgsFirst2, setArgsSecond2 := session.SetArgsForCall(1)
				Expect(setArgsFirst1).To(Equal("authenticated"))
				Expect(setArgsSecond1).To(Equal(true))
				Expect(setArgsFirst2).To(Equal("account_id"))
				Expect(setArgsSecond2).To(Equal(accountID))
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})
		})
	})

	Describe("Logout", func() {
		Context("when attempting to logout", func() {
			It("returns an error when one occurred while retrieving the session", func() {
				store.GetReturns(nil, errors.New("session error"))
				r = httptest.NewRequest("GET", "/v1/logout", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorFetchingSession())

				authController.Logout(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("deletes the session and redirects with no error", func() {
				store.GetReturns(session, nil)
				r = httptest.NewRequest("GET", "/v1/logout", nil)

				authController.Logout(w, r)
				Expect(w.Code).To(Equal(http.StatusFound))
				Expect(session.DeleteCallCount()).To(Equal(1))
				Expect(session.SaveCallCount()).To(Equal(1))
			})
		})
	})
})
