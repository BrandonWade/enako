package models

// RequestPasswordReset a model for RequestPasswordReset requests.
type RequestPasswordReset struct {
	Username string `json:"username" validate:"uname"`
}

// ChangePassword a model for ChangePassword requests.
type ChangePassword struct {
	Token           int64  `json:"token,omitempty"`
	Password        string `json:"password,omitempty" validate:"pword"`
	ConfirmPassword string `json:"confirm_password,omitempty" validate:"pword"`
}
