package models

type APIError interface{}

type apiError struct {
	Errors []error `json:"errors"`
}

func NewAPIError(errors ...error) APIError {
	return &apiError{
		errors,
	}
}
