package helpers

import (
	"fmt"
	"os"
)

var (
	// LoginRoute the route for the login page.
	LoginRoute = fmt.Sprintf("http://%s/login", os.Getenv("API_HOST"))

	// ForgotPasswordRoute the route for the forgot password page.
	ForgotPasswordRoute = fmt.Sprintf("http://%s/password", os.Getenv("API_HOST"))

	// PasswordResetRoute the route for the password reset page.
	PasswordResetRoute = fmt.Sprintf("http://%s/password/reset", os.Getenv("API_HOST"))
)
