package models

// CreateAccount a model for CreateAccount requests.
type CreateAccount struct {
	ID              int64  `json:"id,omitempty"`
	Username        string `json:"username" validate:"uname"`
	Email           string `json:"email,omitempty" validate:"email"`
	Password        string `json:"password,omitempty" validate:"pword"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"pword"`
	ActivationLink  string `json:"activation_link,omitempty"`
}
