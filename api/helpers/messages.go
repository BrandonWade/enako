package helpers

import "fmt"

// MessageActivationEmailSent returned when an account activation email is sent.
func MessageActivationEmailSent(email string) string {
	return fmt.Sprintf("We've sent an email to %s - please follow the instructions inside to activate your account.", email)
}

// MessageActivationEmailAlreadySent returned when an account activation email had previously been sent.
func MessageActivationEmailAlreadySent(email string) string {
	return fmt.Sprintf("Your account has not been activated. We had previously sent an email to %s with the instructions for activating your account. If it isn't in your inbox, please check your spam folder.", email)
}

// MessageResetPasswordEmailSent returned when a reset password request is made.
func MessageResetPasswordEmailSent(email string) string {
	return fmt.Sprintf("Heads up! A password reset link was sent to %s.", email)
}

// MessageAccountWithEmailNotFound returned when an account with the given email wasn't found in the database.
func MessageAccountWithEmailNotFound(email string) string {
	return fmt.Sprintf("Hmm we couldn't find an account with that email: %s. Please make sure that it's spelled correctly and try again.", email)
}

// MessagePasswordUpdated returned when a password reset was successful.
func MessagePasswordUpdated() string {
	return "Hey we were able to successfully update your account! Now please login using your new password."
}
