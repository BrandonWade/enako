package helpers

import (
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	errGeneratingHash = errors.New("error generating password hash")
	errComparingHash  = errors.New("error comparing password and hash")
)

//go:generate counterfeiter -o fakes/fake_password_hasher.go . PasswordHasher
type PasswordHasher interface {
	Generate(password string) (string, error)
	Compare(hash, password string) error
}

type passwordHasher struct {
	logger *logrus.Logger
}

func NewPasswordHasher(logger *logrus.Logger) PasswordHasher {
	return &passwordHasher{
		logger,
	}
}

func (p *passwordHasher) Generate(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":   "PasswordHasher.Generate",
			"password": password,
			"err":      err.Error(),
		}).Error(errGeneratingHash)

		return "", errGeneratingHash
	}

	return string(hash), nil
}

func (p *passwordHasher) Compare(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"method":   "PasswordHasher.Compare",
			"hash":     hash,
			"password": password,
			"err":      err.Error(),
		}).Error(errComparingHash)

		return errComparingHash
	}

	return nil
}
