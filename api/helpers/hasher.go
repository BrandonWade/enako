package helpers

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//go:generate counterfeiter -o fakes/fake_password_hasher.go . PasswordHasher
type PasswordHasher interface {
	Generate(password string) (string, error)
	Compare(hash, password string) error
}

type passwordHasher struct {
	logger *logrus.Logger
}

// NewPasswordHasher ...
func NewPasswordHasher(logger *logrus.Logger) PasswordHasher {
	return &passwordHasher{
		logger,
	}
}

// Generate ...
func (p *passwordHasher) Generate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":   "PasswordHasher.Generate",
			"password": password,
			"err":      err.Error(),
		}).Error(ErrorGeneratingHash())

		return "", ErrorGeneratingHash()
	}

	return string(hash), nil
}

// Compare ...
func (p *passwordHasher) Compare(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":   "PasswordHasher.Compare",
			"hash":     hash,
			"password": password,
			"err":      err.Error(),
		}).Error(ErrorComparingHash())

		return ErrorComparingHash()
	}

	return nil
}
