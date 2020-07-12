package models

// RequestPasswordReset a model for RequestPasswordReset requests.
type RequestPasswordReset struct {
	Email string `json:"email" validate:"email"`
}

// PasswordReset a model for PasswordReset requests.
type PasswordReset struct {
	Password        string `json:"password,omitempty" validate:"pword"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"pword"`
}
