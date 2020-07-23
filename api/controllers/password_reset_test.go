package controllers_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/models"

	"github.com/BrandonWade/enako/api/controllers"
	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/middleware"
	"github.com/BrandonWade/enako/api/services/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("PasswordResetController", func() {
	var (
		logger                  *logrus.Logger
		passwordResetService    *fakes.FakePasswordResetService
		passwordResetController controllers.PasswordResetController
		w                       *httptest.ResponseRecorder
		r                       *http.Request
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		passwordResetService = &fakes.FakePasswordResetService{}
		passwordResetController = controllers.NewPasswordResetController(logger, passwordResetService)

		w = httptest.NewRecorder()
	})

	Describe("RequestPasswordReset", func() {
		Context("when requesting a password reset for an account", func() {
			It("returns an error when one occurs while retrieving the RequestPasswordReset from the request context", func() {
				r = httptest.NewRequest("POST", "/v1/password", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorRequestingPasswordReset())

				passwordResetController.RequestPasswordReset(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns an error if the account with the given email cannot be found", func() {
				accountEmail := "foo@bar.net"

				r = httptest.NewRequest("POST", "/v1/password", nil)
				passwordResetService.RequestPasswordResetReturns("", helpers.ErrorAccountNotFound())
				payload := models.RequestPasswordReset{Email: accountEmail}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"info"}]}`, helpers.MessageAccountWithEmailNotFound(accountEmail))
				ctx := context.WithValue(r.Context(), middleware.ContextRequestPasswordResetKey, payload)
				r = r.WithContext(ctx)

				passwordResetController.RequestPasswordReset(w, r)
				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns an error if some other error occurred when attempting to retrieve the account", func() {
				accountEmail := "foo@bar.net"

				r = httptest.NewRequest("POST", "/v1/password", nil)
				passwordResetService.RequestPasswordResetReturns("", errors.New("service error"))
				payload := models.RequestPasswordReset{Email: accountEmail}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorRequestingPasswordReset())
				ctx := context.WithValue(r.Context(), middleware.ContextRequestPasswordResetKey, payload)
				r = r.WithContext(ctx)

				passwordResetController.RequestPasswordReset(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns a message with no error", func() {
				accountEmail := "foo@bar.net"

				r = httptest.NewRequest("POST", "/v1/password", nil)
				passwordResetService.RequestPasswordResetReturns(accountEmail, nil)
				payload := models.RequestPasswordReset{Email: accountEmail}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"info"}]}`, helpers.MessageResetPasswordEmailSent(accountEmail))
				ctx := context.WithValue(r.Context(), middleware.ContextRequestPasswordResetKey, payload)
				r = r.WithContext(ctx)

				passwordResetController.RequestPasswordReset(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})
		})
	})

	Describe("SetPasswordResetToken", func() {
		Context("when attempting to verify a password reset token and save it to a cookie in the response", func() {
			It("redirects to the login page when a malformed reset token is provided", func() {
				r = httptest.NewRequest("GET", "/v1/password/reset?t=invalidresettoken1234", nil)

				passwordResetController.SetPasswordResetToken(w, r)
				Expect(w.Code).To(Equal(http.StatusSeeOther))
			})

			It("redirects to the login page when an invalid reset token is provided", func() {
				passwordResetService.VerifyPasswordResetTokenReturns(false, errors.New("service error"))
				r = httptest.NewRequest("GET", "/v1/password/reset?t=thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234", nil)

				passwordResetController.SetPasswordResetToken(w, r)
				Expect(w.Code).To(Equal(http.StatusSeeOther))
			})

			It("redirects to the password reset page and sets the given reset token as a cookie in the response", func() {
				token := "thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234"
				passwordResetService.VerifyPasswordResetTokenReturns(true, nil)
				r = httptest.NewRequest("GET", fmt.Sprintf("/v1/password/reset?t=%s", token), nil)

				passwordResetController.SetPasswordResetToken(w, r)
				res := w.Result()
				cookie := res.Cookies()[0]

				Expect(res.StatusCode).To(Equal(http.StatusSeeOther))
				Expect(cookie.Name).To(Equal(controllers.PasswordResetCookieName))
				Expect(cookie.Value).To(Equal(token))
				Expect(cookie.MaxAge).To(Equal(controllers.PasswordResetCookieMaxAge))
			})
		})
	})

	Describe("ResetPassword", func() {
		Context("when attempting to reset the password for an account", func() {
			It("returns an error when attempting to retrieve the cookie containing the password reset token from the request if the cookie does not exist", func() {
				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorRetrievingPasswordReset())

				passwordResetController.ResetPassword(w, r)
				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns an error if one occurred while retrieving the ResetPassword from the request context", func() {
				token := "thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234"

				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				r.Header["Cookie"] = []string{fmt.Sprintf("%s=%s; ", controllers.PasswordResetCookieName, token)}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorRetrievingPasswordReset())

				passwordResetController.ResetPassword(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("redirects to the request password reset page if the password reset token has expired or is invalid", func() {
				token := "thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234"

				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				r.Header["Cookie"] = []string{fmt.Sprintf("%s=%s; ", controllers.PasswordResetCookieName, token)}
				passwordResetService.ResetPasswordReturns(false, helpers.ErrorResetTokenExpiredOrInvalid())

				payload := models.PasswordReset{Password: "testpassword123", ConfirmPassword: "testpassword123"}
				ctx := context.WithValue(r.Context(), middleware.ContextPasswordResetKey, payload)
				r = r.WithContext(ctx)

				passwordResetController.ResetPassword(w, r)
				Expect(w.Code).To(Equal(http.StatusSeeOther))
			})

			It("returns an error if the service returned while resetting the password for the account", func() {
				token := "thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234"

				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				r.Header["Cookie"] = []string{fmt.Sprintf("%s=%s; ", controllers.PasswordResetCookieName, token)}
				passwordResetService.ResetPasswordReturns(false, errors.New("service error"))

				payload := models.PasswordReset{Password: "testpassword123", ConfirmPassword: "testpassword123"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"error"}]}`, helpers.ErrorResettingPassword())
				ctx := context.WithValue(r.Context(), middleware.ContextPasswordResetKey, payload)
				r = r.WithContext(ctx)

				passwordResetController.ResetPassword(w, r)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})

			It("returns a message indicating the password for the account was successfully reset", func() {
				token := "thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234"

				r = httptest.NewRequest("POST", "/v1/password/reset", nil)
				r.Header["Cookie"] = []string{fmt.Sprintf("%s=%s; ", controllers.PasswordResetCookieName, token)}
				passwordResetService.ResetPasswordReturns(true, nil)

				payload := models.PasswordReset{Password: "testpassword123", ConfirmPassword: "testpassword123"}
				resBody := fmt.Sprintf(`{"messages":[{"text":"%s","type":"info"}]}`, helpers.MessagePasswordUpdated())
				ctx := context.WithValue(r.Context(), middleware.ContextPasswordResetKey, payload)
				r = r.WithContext(ctx)

				passwordResetController.ResetPassword(w, r)
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo(string(resBody)))
			})
		})
	})
})
