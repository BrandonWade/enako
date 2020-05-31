package validation

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"time"

	validator "gopkg.in/validator.v2"
)

const (
	minUsernameLength = 5
	maxUsernameLength = 32
	minPasswordLength = 15
	maxPasswordLength = 50
)

var (
	errMustBeString              = errors.New("must be string")
	errUsernameTooShort          = fmt.Errorf("username must be minimum %d characters", minUsernameLength)
	errUsernameTooLong           = fmt.Errorf("username must be maximum %d characters", maxUsernameLength)
	errInvalidUsernameCharacters = errors.New("username may only contain alphanumeric characters and underscores")
	errInvalidEmail              = errors.New("invalid email")
	errPasswordTooShort          = fmt.Errorf("password must be minimum %d characters", minPasswordLength)
	errPasswordTooLong           = fmt.Errorf("password must be maximum %d characters", maxPasswordLength)
	errInvalidPasswordCharacters = errors.New("password may only contain alphanumeric characters and the following symbols: _ ! @ # $ % ^ *")
	errInvalidDate               = errors.New("invalid date")
)

func InitValidator() {
	// Add a username validation rule
	validator.SetValidationFunc("uname", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return errMustBeString
		}

		uname := t.String()
		l := len(uname)

		if l < minUsernameLength {
			return errUsernameTooShort
		}

		if l > maxUsernameLength {
			return errUsernameTooLong
		}

		match, err := regexp.MatchString("^\\w+$", uname)
		if err != nil || match != true {
			return errInvalidUsernameCharacters
		}

		return nil
	})

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
		l := len(pword)

		// Ensure length is compatible with bcrypt requirements
		if l < minPasswordLength {
			return errPasswordTooShort
		}

		if l > maxPasswordLength {
			return errPasswordTooLong
		}

		match, err := regexp.MatchString("^[\\w\\!\\@\\#\\$\\%\\^\\*]+$", pword)
		if err != nil || match != true {
			return errInvalidPasswordCharacters
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
