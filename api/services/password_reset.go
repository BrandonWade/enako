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
	RequestPasswordReset(email string) (string, error)
	CheckPasswordResetTokenIsValid(resetToken *models.PasswordResetToken) (bool, error)
	VerifyPasswordResetToken(token string) (bool, error)
	ResetPassword(token, password string) (bool, error)
}

type passwordResetService struct {
	logger       *logrus.Logger
	hasher       helpers.PasswordHasher
	generator    helpers.TokenGenerator
	emailService EmailService
	repo         repositories.PasswordResetRepository
	accountRepo  repositories.AccountRepository
}

// NewPasswordResetService returns a new instance of an PasswordResetService.
func NewPasswordResetService(logger *logrus.Logger, hasher helpers.PasswordHasher, generator helpers.TokenGenerator, emailService EmailService, repo repositories.PasswordResetRepository, accountRepo repositories.AccountRepository) PasswordResetService {
	return &passwordResetService{
		logger,
		hasher,
		generator,
		emailService,
		repo,
		accountRepo,
	}
}

// RequestPasswordReset requests a password reset for the account with the given email.
func (p *passwordResetService) RequestPasswordReset(email string) (string, error) {
	account, err := p.accountRepo.GetAccountByEmail(email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.logger.WithFields(logrus.Fields{
				"method": "PasswordResetService.RequestPasswordReset",
				"email":  email,
			}).Info(err.Error())

			return "", helpers.ErrorAccountNotFound()
		}

		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.RequestPasswordReset",
			"email":  email,
		}).Error(err.Error())

		return "", helpers.ErrorRequestingPasswordReset()
	}

	token := p.generator.CreateToken(PasswordResetTokenLength)
	_, err = p.repo.CreatePasswordResetToken(account.ID, token)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":    "PasswordResetService.RequestPasswordReset",
			"accountID": account.ID,
			"token":     token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingPasswordReset()
	}

	err = p.emailService.SendPasswordResetEmail(account.Email, token)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.RequestPasswordReset",
			"email":  account.Email,
			"token":  token,
		}).Error(err.Error())
		return "", helpers.ErrorRequestingPasswordReset()
	}

	return email, nil
}

// CheckPasswordResetTokenIsValid checks whether the given password reset token has a status of pending and is not expired.
func (p *passwordResetService) CheckPasswordResetTokenIsValid(resetToken *models.PasswordResetToken) (bool, error) {
	now := time.Now()
	expiresAt, err := time.Parse("2006-01-02 03:04:05", resetToken.ExpiresAt)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":    "PasswordResetService.CheckPasswordResetTokenIsValid",
			"token":     resetToken.ResetToken,
			"expiresAt": resetToken.ExpiresAt,
		}).Error(err.Error())
		return false, err
	}

	if resetToken.Status != "pending" || now.After(expiresAt) {
		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.CheckPasswordResetTokenIsValid",
			"token":  resetToken.ResetToken,
		}).Info(helpers.ErrorResetTokenExpiredOrInvalid())
		return false, helpers.ErrorResetTokenExpiredOrInvalid()
	}

	return true, nil
}

// VerifyPasswordResetToken retrieves the password reset token model using the given token and checks whether it is valid.
func (p *passwordResetService) VerifyPasswordResetToken(token string) (bool, error) {
	resetToken, err := p.repo.GetPasswordResetTokenByPasswordResetToken(token)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.VerifyPasswordResetToken",
			"token":  token,
		}).Error(err.Error())
	}

	return p.CheckPasswordResetTokenIsValid(resetToken)
}

// ResetPassword sets the password for the account associated with the reset token.
func (p *passwordResetService) ResetPassword(token, password string) (bool, error) {
	resetToken, err := p.repo.GetPasswordResetTokenByPasswordResetToken(token)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	_, err = p.CheckPasswordResetTokenIsValid(resetToken)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	hash, err := p.hasher.Generate(password)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":   "PasswordResetService.ResetPassword",
			"password": password,
		}).Error(err.Error())
		return false, err
	}

	_, err = p.repo.ResetPassword(token, string(hash))
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	account, err := p.accountRepo.GetAccountByPasswordResetToken(token)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	err = p.emailService.SendPasswordUpdatedEmail(account.Email)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method": "PasswordResetService.ResetPassword",
			"token":  token,
		}).Error(err.Error())
		return false, err
	}

	return true, nil
}
