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

// InitValidator ...
func InitValidator() {
	// Add a username validation rule (alphanumeric plus underscore)
	validator.SetValidationFunc("uname", func(v interface{}, param string) error {
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
	})

	// Add a simple email validation rule
	validator.SetValidationFunc("email", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return helpers.ErrorMustBeString()
		}

		match, err := regexp.MatchString("^[^@]+@[^\\.@]+\\..+$", t.String())
		if err != nil || match != true {
			return helpers.ErrorInvalidEmail()
		}

		return nil
	})

	// Add a password matching rule (alphanumeric plus symbols)
	validator.SetValidationFunc("pword", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return helpers.ErrorMustBeString()
		}

		pword := t.String()
		l := len(pword)

		// Ensure length is compatible with bcrypt requirements
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
	})

	// Add a date matching rule for checking ISO 8601 dates
	validator.SetValidationFunc("date", func(v interface{}, param string) error {
		t := reflect.ValueOf(v)
		if t.Kind() != reflect.String {
			return helpers.ErrorMustBeString()
		}

		_, err := time.Parse("2006-01-02", t.String())
		if err != nil {
			return helpers.ErrorInvalidDate()
		}

		return nil
	})
}
