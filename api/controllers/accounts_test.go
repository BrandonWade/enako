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
		Context("when registering a new account", func() {
			It("returns an error if one was encountered while communicating with the service", func() {
				accountService.RegisterUserReturns(0, errors.New("service error"))
				r = httptest.NewRequest("POST", "/v1/accounts", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorCreatingAccount())

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
})
