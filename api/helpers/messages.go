package helpers

import "fmt"

// MessageResetPasswordEmailSent returned when a reset password request is made.
func MessageResetPasswordEmailSent(email string) string {
	return fmt.Sprintf("Heads up! A password reset link was sent to the email associated with your account: %s", email)
}

// MessageAccountWithEmailNotFound returned when an account with the given email wasn't found in the database.
func MessageAccountWithEmailNotFound(email string) string {
	return fmt.Sprintf("Hmm we couldn't find an account with that email: %s. Please make sure that it's spelled correctly and try again.", email)
}

// MessagePasswordUpdated returned when a password reset was successful.
func MessagePasswordUpdated() string {
	return fmt.Sprintf("Hey we were able to successfully update your account - please login using your new password.")
}
