package helpers

import "fmt"

// MessageResetPasswordEmailSent returned when a reset password request is made.
func MessageResetPasswordEmailSent(email string) string {
	return fmt.Sprintf("Heads up! A password reset link was sent to the email associated with your account: %s", email)
}

// MessageAccountWithUsernameNotFound returned when an account with the given username wasn't found in the database.
func MessageAccountWithUsernameNotFound(username string) string {
	return fmt.Sprintf("Hmm we couldn't find an account with that username. Please make sure that it's spelled correctly and try again.")
}
