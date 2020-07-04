package helpers

import (
	"errors"
	"fmt"
)

var (
	errInvalidAccountPayload              = errors.New("invalid account payload")
	errRetrievingAccount                  = errors.New("error retrieving account from context")
	errCreatingAccount                    = errors.New("error creating account")
	errPasswordsDoNotMatch                = errors.New("passwords do not match")
	errUserNotAuthenticated               = errors.New("user not authenticated")
	errInvalidExpensePayload              = errors.New("invalid expense payload")
	errRetrievingExpense                  = errors.New("error retrieving expense from context")
	errCreatingExpense                    = errors.New("error creating expense")
	errUpdatingExpense                    = errors.New("error updating expense")
	errFetchingExpenses                   = errors.New("error fetching expense list")
	errInvalidExpenseID                   = errors.New("invalid expense id")
	errNoExpensesUpdated                  = errors.New("no expenses were updated")
	errDeletingExpense                    = errors.New("error deleting expense")
	errNoExpensesDeleted                  = errors.New("no expenses were deleted")
	errFetchingCategories                 = errors.New("error fetching categories")
	errFetchingSession                    = errors.New("error fetching session")
	errInvalidUsernameOrPassword          = errors.New("invalid username or password")
	errMustBeString                       = errors.New("must be string")
	errInvalidUsernameCharacters          = errors.New("username may only contain alphanumeric characters and underscores")
	errInvalidEmail                       = errors.New("invalid email")
	errInvalidPasswordCharacters          = errors.New("password may only contain alphanumeric characters and the following symbols: _ ! @ # $ % ^ *")
	errInvalidDate                        = errors.New("invalid date")
	errRetrievingAccountID                = errors.New("error retrieving user account id")
	errAccountNotActivated                = errors.New("account not activated")
	errInvalidActivationToken             = errors.New("invalid account activation token")
	errActivatingAccount                  = errors.New("error activating account")
	errInvalidRequestPasswordResetPayload = errors.New("invalid request password reset payload")
	errRetrievingRequestPasswordReset     = errors.New("error retrieving request password reset from context")
	errRequestingPasswordReset            = errors.New("error requesting password reset")
	errAccountNotFound                    = errors.New("error account not found")
	errObfuscatingEmail                   = errors.New("error obfuscating email")
	errRetrievingPasswordReset            = errors.New("error retrieving password reset from context")
	errInvalidPasswordResetPayload        = errors.New("invalid password reset payload")
)

// ErrorInvalidAccountPayload returned when an error occurs when a submitted account is malformed.
func ErrorInvalidAccountPayload() error {
	return errInvalidAccountPayload
}

// ErrorRetrievingAccount returned when an error occurs when attempting to retrieve the account model from the request.
func ErrorRetrievingAccount() error {
	return errRetrievingAccount
}

// ErrorCreatingAccount returned when an error occurs when attempting to create an account.
func ErrorCreatingAccount() error {
	return errCreatingAccount
}

// ErrorPasswordsDoNotMatch returned when an error occurs when attempting to create an account and the password and confirm password do not match.
func ErrorPasswordsDoNotMatch() error {
	return errPasswordsDoNotMatch
}

// ErrorUserNotAuthenticated returned when an error occurs when a request hits a secure api endpoint while unauthenticated.
func ErrorUserNotAuthenticated() error {
	return errUserNotAuthenticated
}

// ErrorInvalidExpensePayload returned when an error occurs when a submitted expense is malformed.
func ErrorInvalidExpensePayload() error {
	return errInvalidExpensePayload
}

// ErrorRetrievingExpense returned when an error occurs when attempting to retrieve the expense model from the request.
func ErrorRetrievingExpense() error {
	return errRetrievingExpense
}

// ErrorCreatingExpense returned when an error occurs when attempting to create an expense.
func ErrorCreatingExpense() error {
	return errCreatingExpense
}

// ErrorUpdatingExpense returned when an error occurs when attempting to update an expense.
func ErrorUpdatingExpense() error {
	return errUpdatingExpense
}

// ErrorFetchingExpenses returned when an error occurs when attempting to fetch an expense.
func ErrorFetchingExpenses() error {
	return errFetchingExpenses
}

// ErrorInvalidExpenseID returned when an error occurs when an invalid expense id is submitted.
func ErrorInvalidExpenseID() error {
	return errInvalidExpenseID
}

// ErrorNoExpensesUpdated returned when an error occurs when when no expenses were updated when attempting to update an expense.
func ErrorNoExpensesUpdated() error {
	return errNoExpensesUpdated
}

// ErrorDeletingExpense returned when an error occurs when attempting to delete an expense.
func ErrorDeletingExpense() error {
	return errDeletingExpense
}

// ErrorNoExpensesDeleted returned when an error occurs when no expenses were deleted when attempting to delete an expense.
func ErrorNoExpensesDeleted() error {
	return errNoExpensesDeleted
}

// ErrorFetchingCategories returned when an error occurs when attempting to fetch the categories.
func ErrorFetchingCategories() error {
	return errFetchingCategories
}

// ErrorFetchingSession returned when an error occurs when an error occurred while fetching the session.
func ErrorFetchingSession() error {
	return errFetchingSession
}

// ErrorInvalidUsernameOrPassword returned when an error occurs when an invalid username or password is submitted.
func ErrorInvalidUsernameOrPassword() error {
	return errInvalidUsernameOrPassword
}

// ErrorMustBeString returned when an error occurs when a submitted value is an invalid type.
func ErrorMustBeString() error {
	return errMustBeString
}

// ErrorUsernameTooShort returned when an error occurs when a submitted username is too short.
func ErrorUsernameTooShort(minUsernameLength int) error {
	return fmt.Errorf("username must be minimum %d characters", minUsernameLength)
}

// ErrorUsernameTooLong returned when an error occurs when a submitted username is too long.
func ErrorUsernameTooLong(maxUsernameLength int) error {
	return fmt.Errorf("username must be maximum %d characters", maxUsernameLength)
}

// ErrorInvalidUsernameCharacters returned when an error occurs when a submitted username contains invalid characters.
func ErrorInvalidUsernameCharacters() error {
	return errInvalidUsernameCharacters
}

// ErrorInvalidEmail returned when an error occurs when a submitted email is malformed.
func ErrorInvalidEmail() error {
	return errInvalidEmail
}

// ErrorPasswordTooShort returned when an error occurs when a submitted password is too short.
func ErrorPasswordTooShort(minPasswordLength int) error {
	return fmt.Errorf("password must be minimum %d characters", minPasswordLength)
}

// ErrorPasswordTooLong returned when an error occurs when a submitted password is too long.
func ErrorPasswordTooLong(maxPasswordLength int) error {
	return fmt.Errorf("password must be maximum %d characters", maxPasswordLength)
}

// ErrorInvalidPasswordCharacters returned when an error occurs when a submitted password contains invalid characters.
func ErrorInvalidPasswordCharacters() error {
	return errInvalidPasswordCharacters
}

// ErrorInvalidDate returned when an error occurs when a submitted date is invalid.
func ErrorInvalidDate() error {
	return errInvalidDate
}

// ErrorRetrievingAccountID returned when an error occurs when attempting to retrieve the account ID from the session.
func ErrorRetrievingAccountID() error {
	return errRetrievingAccountID
}

// ErrorAccountNotActivated returned when attempting to log into an account that has not yet been activated.
func ErrorAccountNotActivated() error {
	return errAccountNotActivated
}

// ErrorInvalidActivationToken returned when a problem occurs when attempting to retrieve the activation token from the query string.
func ErrorInvalidActivationToken() error {
	return errInvalidActivationToken
}

// ErrorActivatingAccount returned when a problem occurs when attempting to activate an account.
func ErrorActivatingAccount() error {
	return errActivatingAccount
}

// ErrorInvalidRequestPasswordResetPayload returned when an error occurs when a submitted request password reset request is malformed.
func ErrorInvalidRequestPasswordResetPayload() error {
	return errInvalidRequestPasswordResetPayload
}

// ErrorRetrievingRequestPasswordReset returned when an error occurs when attempting to retrieve the request password reset model from the request.
func ErrorRetrievingRequestPasswordReset() error {
	return errRetrievingRequestPasswordReset
}

// ErrorRequestingPasswordReset returned when an error occurs when attempting to request a password reset.
func ErrorRequestingPasswordReset() error {
	return errRequestingPasswordReset
}

// ErrorAccountNotFound returned when an account is not found in the database.
func ErrorAccountNotFound() error {
	return errAccountNotFound
}

// ErrorObfuscatingEmail returned when an error occurs when attempting to obfuscate an email.
func ErrorObfuscatingEmail() error {
	return errObfuscatingEmail
}

// ErrorInvalidPasswordResetPayload returned when an error occurs when a submitted password reset is malformed.
func ErrorInvalidPasswordResetPayload() error {
	return errInvalidPasswordResetPayload
}

// ErrorRetrievingPasswordReset returned when an error occurs when attempting to retrieve the password reset model from the request.
func ErrorRetrievingPasswordReset() error {
	return errRetrievingPasswordReset
}
