package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/BrandonWade/enako/api/models"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/BrandonWade/enako/api/repositories"

	"github.com/sirupsen/logrus"
)

const (
	// PasswordResetTokenLength the length of a password reset token.
	PasswordResetTokenLength = 64
)

// PasswordResetService an interface for working with password resets.
//go:generate counterfeiter -o fakes/fake_password_reset_service.go . PasswordResetService
type PasswordResetService interface {
	RequestPasswordReset(username string) (string, error)
	CheckPasswordResetTokenIsValid(resetToken *models.PasswordResetToken) error
	VerifyPasswordResetToken(token string) error
	ResetPassword(token, password string) (bool, error)
	NotifyOfPasswordReset(token string) error
}

type passwordResetToken struct {
	logger       *logrus.Logger
	hasher       helpers.PasswordHasher
	obfuscator   helpers.EmailObfuscator
	generator    helpers.TokenGenerator
	emailService EmailService
	repo         repositories.PasswordResetRepository
	accountRepo  repositories.AccountRepository
}

// NewPasswordResetService returns a new instance of an PasswordResetService.
func NewPasswordResetService(logger *logrus.Logger, hasher helpers.PasswordHasher, obfuscator helpers.EmailObfuscator, generator helpers.TokenGenerator, emailService EmailService, repo repositories.PasswordResetRepository, accountRepo repositories.AccountRepository) PasswordResetService {
	return &passwordResetToken{
		logger,
		hasher,
		obfuscator,
		generator,
		emailService,
		repo,
		accountRepo,
	}
}

// RequestPasswordReset requests a password reset for the account with the given username.
func (a *passwordResetToken) RequestPasswordReset(username string) (string, error) {
	account, err := a.accountRepo.GetAccountByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			a.logger.WithFields(logrus.Fields{
				"method":   "PasswordResetService.RequestPasswordReset",
				"username": username,
			}).Info(err.Error())

			return "", helpers.ErrorAccountNotFound()
		}

		a.logger.WithFields(logrus.Fields{
			"method":   "PasswordResetService.RequestPasswordReset",
			"username": username,
		}).Error(err.Error())

		return "", helpers.ErrorRequestingPasswordReset()
	}

	token := a.generator.CreateToken(PasswordResetTokenLength)
	_, err = a.repo.CreatePasswordResetToken(account.ID, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "PasswordResetService.RequestPasswordReset",
			"accountID": account.ID,
			"token":     token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingPasswordReset()
	}

	err = a.emailService.SendPasswordResetEmail(account.Email, token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "PasswordResetService.RequestPasswordReset",
			"username": username,
			"email":    account.Email,
			"token":    token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingPasswordReset()
	}

	email, err := a.obfuscator.Obfuscate(account.Email)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "PasswordResetService.RequestPasswordReset",
			"username": username,
			"email":    account.Email,
		}).Error(err.Error())

		return "", helpers.ErrorRequestingPasswordReset()
	}

	return email, nil
}

// CheckPasswordResetTokenIsValid checks whether the given password reset token has a status of pending and is not expired.
func (a *passwordResetToken) CheckPasswordResetTokenIsValid(resetToken *models.PasswordResetToken) error {
	now := time.Now()
	expiresAt, err := time.Parse("2006-01-02 03:04:05", resetToken.ExpiresAt)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":    "PasswordResetService.CheckPasswordResetTokenIsValid",
			"token":     resetToken.ResetToken,
			"expiresAt": resetToken.ExpiresAt,
		}).Error(err.Error())
		return err
	}

	if resetToken.Status != "pending" || now.After(expiresAt) {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.CheckPasswordResetTokenIsValid",
			"token":  resetToken.ResetToken,
		}).Info(helpers.ErrorResetTokenExpiredOrInvalid())
		return helpers.ErrorResetTokenExpiredOrInvalid()
	}

	return nil
}

// VerifyPasswordResetToken retrieves the password reset token model using the given token and checks whether it is valid.
func (a *passwordResetToken) VerifyPasswordResetToken(token string) error {
	resetToken, err := a.repo.GetPasswordResetTokenByPasswordResetToken(token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.VerifyPasswordResetToken",
			"token":  token,
		}).Error(err.Error())
	}

	return a.CheckPasswordResetTokenIsValid(resetToken)
}

// ResetPassword sets the password for the account associated with the reset token.
func (a *passwordResetToken) ResetPassword(token, password string) (bool, error) {
	resetToken, err := a.repo.GetPasswordResetTokenByPasswordResetToken(token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	err = a.CheckPasswordResetTokenIsValid(resetToken)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	hash, err := a.hasher.Generate(password)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method":   "PasswordResetService.ResetPassword",
			"password": password,
		}).Error(err.Error())
		return false, err
	}

	_, err = a.repo.ResetPassword(token, string(hash))
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	err = a.NotifyOfPasswordReset(token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	return true, nil
}

// NotifyOfPasswordReset notifies the account owner that their password was reset.
func (a *passwordResetToken) NotifyOfPasswordReset(token string) error {
	account, err := a.accountRepo.GetAccountByPasswordResetToken(token)
	if err != nil {
		a.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.NotifyOfPasswordReset",
			"token":  token,
		}).Error(err.Error())
		return err
	}

	return a.emailService.SendPasswordUpdatedEmail(account.Email)
}
