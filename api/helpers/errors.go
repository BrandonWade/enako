package helpers

import (
	"errors"
	"fmt"
)

// ErrorInvalidAccountPayload ...
func ErrorInvalidAccountPayload() error {
	return errors.New("invalid account payload")
}

// ErrorRetrievingAccount ...
func ErrorRetrievingAccount() error {
	return errors.New("error retrieving account from context")
}

// ErrorCreatingAccount ...
func ErrorCreatingAccount() error {
	return errors.New("error creating account")
}

// ErrorVerifyingAccount ...
func ErrorVerifyingAccount() error {
	return errors.New("error verifying account")
}

// ErrorPasswordsDoNotMatch ...
func ErrorPasswordsDoNotMatch() error {
	return errors.New("passwords do not match")
}

// ErrorUserNotAuthenticated ...
func ErrorUserNotAuthenticated() error {
	return errors.New("user not authenticated")
}

// ErrorInvalidExpensePayload ...
func ErrorInvalidExpensePayload() error {
	return errors.New("invalid expense payload")
}

// ErrorRetrievingExpense ...
func ErrorRetrievingExpense() error {
	return errors.New("error retrieving expense from context")
}

// ErrorCreatingExpense ...
func ErrorCreatingExpense() error {
	return errors.New("error creating expense")
}

// ErrorUpdatingExpense ...
func ErrorUpdatingExpense() error {
	return errors.New("error updating expense")
}

// ErrorFetchingExpenses ...
func ErrorFetchingExpenses() error {
	return errors.New("error fetching expense list")
}

// ErrorInvalidExpenseID ...
func ErrorInvalidExpenseID() error {
	return errors.New("invalid expense id")
}

// ErrorNoExpensesUpdated ...
func ErrorNoExpensesUpdated() error {
	return errors.New("no expenses were updated")
}

// ErrorDeletingExpense ...
func ErrorDeletingExpense() error {
	return errors.New("error deleting expense")
}

// ErrorNoExpensesDeleted ...
func ErrorNoExpensesDeleted() error {
	return errors.New("no expenses were deleted")
}

// ErrorFetchingCategories ...
func ErrorFetchingCategories() error {
	return errors.New("error fetching categories")
}

// ErrorFetchingSession ...
func ErrorFetchingSession() error {
	return errors.New("error fetching session")
}

// ErrorGeneratingHash ...
func ErrorGeneratingHash() error {
	return errors.New("error generating password hash")
}

// ErrorComparingHash ...
func ErrorComparingHash() error {
	return errors.New("error comparing password and hash")
}

// ErrorInvalidUsernameOrPassword ...
func ErrorInvalidUsernameOrPassword() error {
	return errors.New("invalid username or password")
}

// ErrorMustBeString ...
func ErrorMustBeString() error {
	return errors.New("must be string")
}

// ErrorUsernameTooShort ...
func ErrorUsernameTooShort(minUsernameLength int) error {
	return fmt.Errorf("username must be minimum %d characters", minUsernameLength)
}

// ErrorUsernameTooLong ...
func ErrorUsernameTooLong(maxUsernameLength int) error {
	return fmt.Errorf("username must be maximum %d characters", maxUsernameLength)
}

// ErrorInvalidUsernameCharacters ...
func ErrorInvalidUsernameCharacters() error {
	return errors.New("username may only contain alphanumeric characters and underscores")
}

// ErrorInvalidEmail ...
func ErrorInvalidEmail() error {
	return errors.New("invalid email")
}

// ErrorPasswordTooShort ...
func ErrorPasswordTooShort(minPasswordLength int) error {
	return fmt.Errorf("password must be minimum %d characters", minPasswordLength)
}

// ErrorPasswordTooLong ...
func ErrorPasswordTooLong(maxPasswordLength int) error {
	return fmt.Errorf("password must be maximum %d characters", maxPasswordLength)
}

// ErrorInvalidPasswordCharacters ...
func ErrorInvalidPasswordCharacters() error {
	return errors.New("password may only contain alphanumeric characters and the following symbols: _ ! @ # $ % ^ *")
}

// ErrorInvalidDate ...
func ErrorInvalidDate() error {
	return errors.New("invalid date")
}
