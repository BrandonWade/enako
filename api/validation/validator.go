package validation

import (
	"reflect"
	"regexp"
	"time"

	"github.com/BrandonWade/enako/api/helpers"
	validator "gopkg.in/validator.v2"
)

const (
	minUsernameLength = 5
	maxUsernameLength = 32
	minPasswordLength = 15
	maxPasswordLength = 50
)

// InitValidator registers custom validation rules with the validator.
func InitValidator() {
	validator.SetValidationFunc("uname", ValidateUsername)
	validator.SetValidationFunc("email", ValidateEmail)
	validator.SetValidationFunc("pword", ValidatePassword)
	validator.SetValidationFunc("date", ValidateDate)
}

// ValidateUsername checks that a username is valid (alphanumeric plus underscores).
func ValidateUsername(v interface{}, param string) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.String {
		return helpers.ErrorMustBeString()
	}

	uname := t.String()
	l := len(uname)

	if l < minUsernameLength {
		return helpers.ErrorUsernameTooShort(minUsernameLength)
	}

	if l > maxUsernameLength {
		return helpers.ErrorUsernameTooLong(maxUsernameLength)
	}

	match, err := regexp.MatchString("^\\w+$", uname)
	if err != nil || match != true {
		return helpers.ErrorInvalidUsernameCharacters()
	}

	return nil
}

// ValidateEmail checks that an email is valid.
func ValidateEmail(v interface{}, param string) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.String {
		return helpers.ErrorMustBeString()
	}

	match, err := regexp.MatchString("^[^@]+@[^\\.@]+\\..+$", t.String())
	if err != nil || match != true {
		return helpers.ErrorInvalidEmail()
	}

	return nil
}

// ValidatePassword checks that a password is valid (alphanumeric plus symbols).
func ValidatePassword(v interface{}, param string) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.String {
		return helpers.ErrorMustBeString()
	}

	pword := t.String()
	l := len(pword)

	// Ensure length is compatible with bcrypt requirements.
	if l < minPasswordLength {
		return helpers.ErrorPasswordTooShort(minPasswordLength)
	}

	if l > maxPasswordLength {
		return helpers.ErrorPasswordTooLong(maxPasswordLength)
	}

	match, err := regexp.MatchString("^[\\w\\!\\@\\#\\$\\%\\^\\*]+$", pword)
	if err != nil || match != true {
		return helpers.ErrorInvalidPasswordCharacters()
	}

	return nil
}

// ValidateDate checks that a date is valid (dates are in ISO 8601 format).
func ValidateDate(v interface{}, param string) error {
	t := reflect.ValueOf(v)
	if t.Kind() != reflect.String {
		return helpers.ErrorMustBeString()
	}

	_, err := time.Parse("2006-01-02", t.String())
	if err != nil {
		return helpers.ErrorInvalidDate()
	}

	return nil
}
