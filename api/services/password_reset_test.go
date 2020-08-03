package services_test

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"time"

	"github.com/BrandonWade/enako/api/helpers"
	helperfakes "github.com/BrandonWade/enako/api/helpers/fakes"
	"github.com/BrandonWade/enako/api/models"
	repositoryfakes "github.com/BrandonWade/enako/api/repositories/fakes"
	"github.com/BrandonWade/enako/api/services"
	servicefakes "github.com/BrandonWade/enako/api/services/fakes"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PasswordResetService", func() {
	var (
		logger               *logrus.Logger
		hasher               *helperfakes.FakePasswordHasher
		generator            *helperfakes.FakeTokenGenerator
		emailService         *servicefakes.FakeEmailService
		accountRepo          *repositoryfakes.FakeAccountRepository
		passwordResetRepo    *repositoryfakes.FakePasswordResetRepository
		passwordResetService services.PasswordResetService
		accountID            int64
		token                string
		email                string
		password             string
		hashedPassword       string
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		hasher = &helperfakes.FakePasswordHasher{}
		generator = &helperfakes.FakeTokenGenerator{}
		emailService = &servicefakes.FakeEmailService{}
		accountRepo = &repositoryfakes.FakeAccountRepository{}
		passwordResetRepo = &repositoryfakes.FakePasswordResetRepository{}
		passwordResetService = services.NewPasswordResetService(logger, hasher, generator, emailService, passwordResetRepo, accountRepo)

		accountID = int64(18742356)
		token = "thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234"
		email = "foo@bar.net"
		password = "testpassword123"
		hashedPassword = "hashedtestpassword123"
	})

	Describe("RequestPasswordReset", func() {
		Context("when requesting a password reset", func() {
			It("returns an error when no account with the given email could be found", func() {
				accountRepo.GetAccountByEmailReturns(&models.Account{}, sql.ErrNoRows)

				email, err := passwordResetService.RequestPasswordReset(email)
				Expect(email).To(BeEmpty())
				Expect(err).To(Equal(helpers.ErrorAccountNotFound()))
			})

			It("returns an error if one occurred while retrieving the account from the repository", func() {
				accountRepo.GetAccountByEmailReturns(&models.Account{}, errors.New("repo error"))

				email, err := passwordResetService.RequestPasswordReset(email)
				Expect(email).To(BeEmpty())
				Expect(err).To(Equal(helpers.ErrorRequestingPasswordReset()))
			})

			It("returns an error if one occurred while creating a password reset token", func() {
				account := &models.Account{ID: accountID, Email: email, Password: password, IsActivated: true}
				accountRepo.GetAccountByEmailReturns(account, nil)
				passwordResetRepo.CreatePasswordResetTokenReturns(0, errors.New("repo error"))

				email, err := passwordResetService.RequestPasswordReset(email)
				Expect(email).To(BeEmpty())
				Expect(err).To(Equal(helpers.ErrorRequestingPasswordReset()))
			})

			It("returns an error if one occurred while sending a password reset email", func() {
				account := &models.Account{ID: accountID, Email: email, Password: password, IsActivated: true}
				accountRepo.GetAccountByEmailReturns(account, nil)
				passwordResetRepo.CreatePasswordResetTokenReturns(1, nil)
				emailService.SendPasswordResetEmailReturns(errors.New("email error"))

				email, err := passwordResetService.RequestPasswordReset(email)
				Expect(email).To(BeEmpty())
				Expect(err).To(Equal(helpers.ErrorRequestingPasswordReset()))
			})

			It("returns the email and no error", func() {
				account := &models.Account{ID: accountID, Email: email, Password: password, IsActivated: true}
				accountRepo.GetAccountByEmailReturns(account, nil)
				passwordResetRepo.CreatePasswordResetTokenReturns(1, nil)
				emailService.SendPasswordResetEmailReturns(nil)

				email, err := passwordResetService.RequestPasswordReset(email)
				Expect(email).To(Equal(email))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("CheckPasswordResetTokenIsValid", func() {
		Context("when checking that a password reset token is valid", func() {
			It("returns an error if one occurred while parsing the expires at time", func() {
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "used", ExpiresAt: ""}

				success, err := passwordResetService.CheckPasswordResetTokenIsValid(resetToken)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if the password reset token is used, disabled, or expired", func() {
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "used", ExpiresAt: "2020-01-01 00:00:00"}

				success, err := passwordResetService.CheckPasswordResetTokenIsValid(resetToken)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns no error", func() {
				expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "pending", ExpiresAt: expiresAt}

				success, err := passwordResetService.CheckPasswordResetTokenIsValid(resetToken)
				Expect(success).To(BeTrue())
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("VerifyPasswordResetToken", func() {
		Context("when verifying a password reset token", func() {
			It("returns an error if one occurred while attempting to retrieve the corresponding password reset token model", func() {
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(&models.PasswordResetToken{}, errors.New("password repo error"))

				success, err := passwordResetService.VerifyPasswordResetToken(token)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns no error", func() {
				expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "pending", ExpiresAt: expiresAt}
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(resetToken, nil)

				success, err := passwordResetService.VerifyPasswordResetToken(token)
				Expect(success).To(BeTrue())
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("ResetPassword", func() {
		Context("when resetting a password", func() {
			It("returns an error if one occurred when attempting to retrieve the corresponding password reset token model", func() {
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(&models.PasswordResetToken{}, errors.New("repo error"))

				success, err := passwordResetService.ResetPassword(token, password)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while checking if the password reset token is valid", func() {
				expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "used", ExpiresAt: expiresAt}
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(resetToken, nil)

				success, err := passwordResetService.ResetPassword(token, password)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while hashing the new password", func() {
				expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "pending", ExpiresAt: expiresAt}
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(resetToken, nil)
				hasher.GenerateReturns("", errors.New("hasher error"))

				success, err := passwordResetService.ResetPassword(token, password)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while updating the account with the new password", func() {
				expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "pending", ExpiresAt: expiresAt}
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(resetToken, nil)
				hasher.GenerateReturns(hashedPassword, nil)
				passwordResetRepo.ResetPasswordReturns(false, errors.New("repo error"))

				success, err := passwordResetService.ResetPassword(token, password)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while retrieving the account associated with the password reset token", func() {
				expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "pending", ExpiresAt: expiresAt}
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(resetToken, nil)
				hasher.GenerateReturns(hashedPassword, nil)
				passwordResetRepo.ResetPasswordReturns(true, nil)
				accountRepo.GetAccountByPasswordResetTokenReturns(&models.Account{}, errors.New("repo error"))

				success, err := passwordResetService.ResetPassword(token, password)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while sending a password updated notification email", func() {
				expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "pending", ExpiresAt: expiresAt}
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(resetToken, nil)
				hasher.GenerateReturns(hashedPassword, nil)
				passwordResetRepo.ResetPasswordReturns(true, nil)
				account := &models.Account{ID: accountID, Email: email, Password: password, IsActivated: true}
				accountRepo.GetAccountByPasswordResetTokenReturns(account, nil)
				emailService.SendPasswordUpdatedEmailReturns(errors.New("email error"))

				success, err := passwordResetService.ResetPassword(token, password)
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})

			It("returns no error", func() {
				expiresAt := time.Now().Add(24 * time.Hour).Format("2006-01-02 15:04:05")
				resetToken := &models.PasswordResetToken{ID: 1, AccountID: accountID, ResetToken: token, Status: "pending", ExpiresAt: expiresAt}
				passwordResetRepo.GetPasswordResetTokenByPasswordResetTokenReturns(resetToken, nil)
				hasher.GenerateReturns(hashedPassword, nil)
				passwordResetRepo.ResetPasswordReturns(true, nil)
				account := &models.Account{ID: accountID, Email: email, Password: password, IsActivated: true}
				accountRepo.GetAccountByPasswordResetTokenReturns(account, nil)
				emailService.SendPasswordUpdatedEmailReturns(nil)

				success, err := passwordResetService.ResetPassword(token, password)
				Expect(success).To(BeTrue())
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
