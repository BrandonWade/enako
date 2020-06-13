package helpers

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher an interface for securely hashing passwords.
//go:generate counterfeiter -o fakes/fake_password_hasher.go . PasswordHasher
type PasswordHasher interface {
	Generate(password string) (string, error)
	Compare(hash, password string) error
}

type passwordHasher struct {
	logger *logrus.Logger
}

// NewPasswordHasher returns a new instance of a passwordHasher.
func NewPasswordHasher(logger *logrus.Logger) PasswordHasher {
	return &passwordHasher{
		logger,
	}
}

// Generate returns a hash for the provided password.
func (p *passwordHasher) Generate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":   "PasswordHasher.Generate",
			"password": password,
		}).Error(err.Error())
		return "", err
	}

	return string(hash), nil
}

// Compare compares a hashed and unhashed password.
func (p *passwordHasher) Compare(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":   "PasswordHasher.Compare",
			"hash":     hash,
			"password": password,
		}).Error(err.Error())
		return err
	}

	return nil
}
