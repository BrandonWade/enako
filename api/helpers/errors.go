package helpers

import (
	"errors"
	"fmt"
)

// ErrorInvalidAccountPayload returned when an error occurs when a submitted account is malformed.
func ErrorInvalidAccountPayload() error {
	return errors.New("invalid account payload")
}

// ErrorRetrievingAccount returned when an error occurs when attempting to retrieve the account from the request.
func ErrorRetrievingAccount() error {
	return errors.New("error retrieving account from context")
}

// ErrorCreatingAccount returned when an error occurs when attempting to create an account.
func ErrorCreatingAccount() error {
	return errors.New("error creating account")
}

// ErrorPasswordsDoNotMatch returned when an error occurs when attempting to create an account and the password and confirm password do not match.
func ErrorPasswordsDoNotMatch() error {
	return errors.New("passwords do not match")
}

// ErrorUserNotAuthenticated returned when an error occurs when a request hits a secure api endpoint while unauthenticated.
func ErrorUserNotAuthenticated() error {
	return errors.New("user not authenticated")
}

// ErrorInvalidExpensePayload returned when an error occurs when a submitted expense is malformed.
func ErrorInvalidExpensePayload() error {
	return errors.New("invalid expense payload")
}

// ErrorRetrievingExpense returned when an error occurs when attempting to retrieve the expense from the request.
func ErrorRetrievingExpense() error {
	return errors.New("error retrieving expense from context")
}

// ErrorCreatingExpense returned when an error occurs when attempting to create an expense.
func ErrorCreatingExpense() error {
	return errors.New("error creating expense")
}

// ErrorUpdatingExpense returned when an error occurs when attempting to update an expense.
func ErrorUpdatingExpense() error {
	return errors.New("error updating expense")
}

// ErrorFetchingExpenses returned when an error occurs when attempting to fetch an expense.
func ErrorFetchingExpenses() error {
	return errors.New("error fetching expense list")
}

// ErrorInvalidExpenseID returned when an error occurs when an invalid expense id is submitted.
func ErrorInvalidExpenseID() error {
	return errors.New("invalid expense id")
}

// ErrorNoExpensesUpdated returned when an error occurs when when no expenses were updated when attempting to update an expense.
func ErrorNoExpensesUpdated() error {
	return errors.New("no expenses were updated")
}

// ErrorDeletingExpense returned when an error occurs when attempting to delete an expense.
func ErrorDeletingExpense() error {
	return errors.New("error deleting expense")
}

// ErrorNoExpensesDeleted returned when an error occurs when no expenses were deleted when attempting to delete an expense.
func ErrorNoExpensesDeleted() error {
	return errors.New("no expenses were deleted")
}

// ErrorFetchingCategories returned when an error occurs when attempting to fetch the categories.
func ErrorFetchingCategories() error {
	return errors.New("error fetching categories")
}

// ErrorFetchingSession returned when an error occurs when an error occurred while fetching the session.
func ErrorFetchingSession() error {
	return errors.New("error fetching session")
}

// ErrorInvalidUsernameOrPassword returned when an error occurs when an invalid username or password is submitted.
func ErrorInvalidUsernameOrPassword() error {
	return errors.New("invalid username or password")
}

// ErrorMustBeString returned when an error occurs when a submitted value is an invalid type.
func ErrorMustBeString() error {
	return errors.New("must be string")
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
	return errors.New("username may only contain alphanumeric characters and underscores")
}

// ErrorInvalidEmail returned when an error occurs when a submitted email is malformed.
func ErrorInvalidEmail() error {
	return errors.New("invalid email")
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
	return errors.New("password may only contain alphanumeric characters and the following symbols: _ ! @ # $ % ^ *")
}

// ErrorInvalidDate returned when an error occurs when a submitted date is invalid.
func ErrorInvalidDate() error {
	return errors.New("invalid date")
}

// ErrorRetrievingUserAccountID returned when an error occurs when attempting to retrieve the user account ID from the session.
func ErrorRetrievingUserAccountID() error {
	return errors.New("error retrieving user account id")
}
