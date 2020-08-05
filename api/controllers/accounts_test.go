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
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("AccountController", func() {
	var (
		logger            *logrus.Logger
		accountService    *fakes.FakeAccountService
		accountController controllers.AccountController
		w                 *httptest.ResponseRecorder
		r                 *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		accountService = &fakes.FakeAccountService{}
		accountController = controllers.NewAccountController(logger, accountService)

		w = httptest.NewRecorder()
	})

	Describe("RegisterUser", func() {
		Context("when attempting to register a new account", func() {
			It("returns an error if one was encountered while retrieving the CreateAccount from the request context", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorCreatingAccount())

				accountController.RegisterUser(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if one was encountered while communicating with the service", func() {
				accountEmail := "email@test.com"

				accountService.RegisterUserReturns(0, errors.New("service error"))
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				payload := models.CreateAccount{Email: accountEmail, Password: "testpassword123", ConfirmPassword: "testpassword123"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorCreatingAccount())
				ctx := context.WithValue(r.Context(), middleware.ContextCreateAccountKey, payload)
				r = r.WithContext(ctx)

				accountController.RegisterUser(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns the info for the created account with no error", func() {
				accountID := int64(100)
				accountEmail := "email@test.com"

				accountService.RegisterUserReturns(accountID, nil)
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				payload := models.CreateAccount{Email: accountEmail, Password: "testpassword123", ConfirmPassword: "testpassword123"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"info"}]}`, helpers.MessageActivationEmailSent(accountEmail))
				ctx := context.WithValue(r.Context(), middleware.ContextCreateAccountKey, payload)
				r = r.WithContext(ctx)

				accountController.RegisterUser(w, r)
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})
		})
	})

	Describe("ActivateAccount", func() {
		Context("when attempting to activate an account", func() {
			It("returns an error if the activation token isn't the proper length", func() {
				r = httptest.NewRequest("GET", "/v1/accounts/activate?t=abc123", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorActivatingAccount())

				accountController.ActivateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("will return an error if one is encountered while communicating with the accounts service", func() {
				accountService.ActivateAccountReturns(false, errors.New("service error"))
				r = httptest.NewRequest("GET", "/v1/accounts/activate?t=thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorActivatingAccount())

				accountController.ActivateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("will redirect to the login page if the account was successfully activated", func() {
				accountService.ActivateAccountReturns(true, nil)
				r = httptest.NewRequest("GET", "/v1/accounts/activate?t=thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234", nil)

				accountController.ActivateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusSeeOther))
			})
		})
	})

	Describe("ChangePassword", func() {
		Context("when attempting to change the password for an account", func() {
			It("returns an error if one was encountered while retrieving the account ID from the request context", func() {
				r = httptest.NewRequest("PUT", "/v1/accounts/password/change", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorChangingPassword())

				accountController.ChangePassword(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if one was encountered while retrieving the ChangePassword from the request context", func() {
				r = httptest.NewRequest("PUT", "/v1/accounts/password/change", nil)
				accountID := int64(12345)
				ctx := context.WithValue(r.Context(), middleware.ContextAccountIDKey, accountID)
				r = r.WithContext(ctx)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorChangingPassword())

				accountController.ChangePassword(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns an error if one was encountered while communicating with the service", func() {
				accountService.ChangePasswordReturns(false, errors.New("service error"))
				r = httptest.NewRequest("PUT", "/v1/accounts/password/change", nil)
				accountID := int64(12345)
				ctx := context.WithValue(r.Context(), middleware.ContextAccountIDKey, accountID)
				r = r.WithContext(ctx)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "testpassword12345", ConfirmPassword: "testpassword12345"}
				ctx = context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorChangingPassword())

				accountController.ChangePassword(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})

			It("returns a message indicating the password was updated and no error", func() {
				accountService.ChangePasswordReturns(true, nil)
				r = httptest.NewRequest("PUT", "/v1/accounts/password/change", nil)
				accountID := int64(12345)
				ctx := context.WithValue(r.Context(), middleware.ContextAccountIDKey, accountID)
				r = r.WithContext(ctx)
				payload := models.ChangePassword{CurrentPassword: "testpassword123", NewPassword: "testpassword12345", ConfirmPassword: "testpassword12345"}
				ctx = context.WithValue(r.Context(), middleware.ContextChangePasswordKey, payload)
				r = r.WithContext(ctx)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"info"}]}`, helpers.MessagePasswordUpdated())

				accountController.ChangePassword(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(resBody))
			})
		})
	})
})
