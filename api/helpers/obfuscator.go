package helpers

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// EmailObfuscator an interface for obfuscating email addresses.
//go:generate counterfeiter -o fakes/fake_email_obfuscator.go . EmailObfuscator
type EmailObfuscator interface {
	Obfuscate(email string) (string, error)
}

type emailObfuscator struct {
	logger *logrus.Logger
}

// NewEmailObfuscator returns a new instance of an EmailObfuscator.
func NewEmailObfuscator(logger *logrus.Logger) EmailObfuscator {
	return &emailObfuscator{
		logger,
	}
}

// Obfuscate obfuscates the provided email address.
func (e *emailObfuscator) Obfuscate(email string) (string, error) {
	at := strings.Index(email, "@")
	if at == -1 {
		e.logger.WithFields(logrus.Fields{
			"method": "EmailObfuscator.Obfuscate",
			"email":  email,
		}).Error(ErrorObfuscatingEmail())

		return "", ErrorObfuscatingEmail()
	}

	username := email[:at]
	l := len(username)

	index := 0
	switch {
	case l == 2:
		fallthrough
	case l == 3:
		index = 1
	case l > 3:
		index = 2
	}

	start := email[:index]
	rest := email[index:at]

	obfUsername := fmt.Sprintf("%s%s", start, strings.Repeat("*", len(rest)))
	obfEmail := strings.Replace(email, username, obfUsername, 1)

	return obfEmail, nil
}
