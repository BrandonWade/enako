package helpers

import "errors"

// ErrorFetchingSession ...
func ErrorFetchingSession() error {
	return errors.New("error fetching session")
}

// ErrorUserNotAuthenticated ...
func ErrorUserNotAuthenticated() error {
	return errors.New("user not authenticated")
}
