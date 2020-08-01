package models

// ChangePassword a model for ChangePassword requests.
type ChangePassword struct {
	CurrentPassword string `json:"current_password,omitempty" validate:"pword"`
	NewPassword     string `json:"new_password,omitempty" validate:"pword"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"pword"`
}
