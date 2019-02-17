package validation

import (
	"errors"
	"reflect"
	"regexp"
	"time"

	validator "gopkg.in/validator.v2"
)

const (
	minPasswordLength = 15
	maxPasswordLength = 50
)

var (
	errMustBeString    = errors.New("must be string")
	errInvalidEmail    = errors.New("invalid email")
	errInvalidPassword = errors.New("invalid password")
	errInvalidDate     = errors.New("invalid date")
)

func InitValidator() {
	// Add a simple email validation rule
	validator.SetValidationFunc("email", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return errMustBeString
		}

		match, err := regexp.MatchString("^[^@]+@[^\\.@]+\\..+$", t.String())
		if err != nil || match != true {
			return errInvalidEmail
		}

		return nil
	})

	// Add a password matching rule (alphanumeric plus symbols)
	validator.SetValidationFunc("pword", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return errMustBeString
		}

		pword := t.String()

		// Ensure length is compatible with bcrypt requirements
		l := len(pword)
		if l < minPasswordLength || l > maxPasswordLength {
			return errInvalidPassword
		}

		match, err := regexp.MatchString("^[\\w\\!\\@\\#\\$\\%\\^\\&\\*]+$", pword)
		if err != nil || match != true {
			return errInvalidPassword
		}

		return nil
	})

	// Add a date matching rule for checking ISO 8601 dates
	validator.SetValidationFunc("date", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return errMustBeString
		}

		_, err := time.Parse("2006-01-02", t.String())
		if err != nil {
			return errInvalidDate
		}

		return nil
	})
}
