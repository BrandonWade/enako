package helpers

import "errors"

// ErrorInvalidAccountPayload ...
func ErrorInvalidAccountPayload() error {
	return errors.New("invalid account payload")
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

// ErrorFetchingSession ...
func ErrorFetchingSession() error {
	return errors.New("error fetching session")
}
