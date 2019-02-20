package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/services/fakes"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthController", func() {
	var (
		logger         *logrus.Logger
		authService    *fakes.FakeAuthService
		authController controllers.AuthController
		w              *httptest.ResponseRecorder
		r              *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		authService = &fakes.FakeAuthService{}
		authController = controllers.NewAuthController(logger, authService)

		w = httptest.NewRecorder()
	})

	Describe("CreateAccount", func() {

		Context("when creating a new account", func() {

			It("returns an error if a malformed payload is submitted", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", strings.NewReader(`{foo}`))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrInvalidAccountPayload)

				authController.CreateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if an invalid email is submitted", func() {
				payload := models.UserAccount{UserAccountEmail: "invalid@@invalid.com", UserAccountPassword: "testpassword123", ConfirmPassword: "testpassword123"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))
				resBody := `{"errors":["UserAccountEmail: invalid email"]}`

				authController.CreateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if a password that is too short is submitted", func() {
				payload := models.UserAccount{UserAccountEmail: "test@email.com", UserAccountPassword: "password", ConfirmPassword: "password"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))
				invalidPword := `{"errors":["UserAccountPassword: invalid password"]}`
				invalidConfirmPword := `{"errors":["ConfirmPassword: invalid password"]}`

				authController.CreateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(Or(BeEquivalentTo(invalidPword+"\n"), BeEquivalentTo(invalidConfirmPword+"\n")))
			})

			It("returns an error if a password that is too long is submitted", func() {
				payload := models.UserAccount{UserAccountEmail: "test@email.com", UserAccountPassword: "thisisareallylongpasswordthatistoolongandwillfailvalidation", ConfirmPassword: "thisisareallylongpasswordthatistoolongandwillfailvalidation"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))
				invalidPword := `{"errors":["UserAccountPassword: invalid password"]}`
				invalidConfirmPword := `{"errors":["ConfirmPassword: invalid password"]}`

				authController.CreateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(Or(BeEquivalentTo(invalidPword+"\n"), BeEquivalentTo(invalidConfirmPword+"\n")))
			})

			It("returns an error if a password that contains invalid characters is submitted", func() {
				payload := models.UserAccount{UserAccountEmail: "test@email.com", UserAccountPassword: "_-123testpassword456-_", ConfirmPassword: "_-123testpassword456-_"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))
				invalidPword := `{"errors":["UserAccountPassword: invalid password"]}`
				invalidConfirmPword := `{"errors":["ConfirmPassword: invalid password"]}`

				authController.CreateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(Or(BeEquivalentTo(invalidPword+"\n"), BeEquivalentTo(invalidConfirmPword+"\n")))
			})

			It("returns an error if the passwords do not match", func() {
				payload := models.UserAccount{UserAccountEmail: "email@test.com", UserAccountPassword: "testpassword123", ConfirmPassword: "testpassword1234"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrPasswordsDoNotMatch)

				authController.CreateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusUnprocessableEntity))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns an error if one was encountered while communicating with the service", func() {
				authService.CreateAccountReturns(0, errors.New("service error"))
				payload := models.UserAccount{UserAccountEmail: "email@test.com", UserAccountPassword: "testpassword123", ConfirmPassword: "testpassword123"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))
				resBody := fmt.Sprintf(`{"errors":["%s"]}`, controllers.ErrCreatingAccount)

				authController.CreateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(w.Body.String()).To(BeEquivalentTo(resBody + "\n"))
			})

			It("returns the info for the created account with no error", func() {
				accountID := int64(100)
				accountEmail := "email@test.com"

				authService.CreateAccountReturns(accountID, nil)
				payload := models.UserAccount{UserAccountEmail: accountEmail, UserAccountPassword: "testpassword123", ConfirmPassword: "testpassword123"}
				payloadJSON, _ := json.Marshal(payload)

				r = httptest.NewRequest("POST", "/v1/accounts", bytes.NewBuffer(payloadJSON))

				response := models.UserAccount{ID: accountID, UserAccountEmail: accountEmail}
				responseJSON, _ := json.Marshal(response)

				authController.CreateAccount(w, r)
				Expect(w.Code).To(Equal(http.StatusCreated))
				Expect(w.Body.String()).To(BeEquivalentTo(string(responseJSON) + "\n"))
			})
		})
	})
})
