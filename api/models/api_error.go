package models

// APIError a wrapper interface for formatting error messages.
type APIError interface{}

type apiError struct {
	Errors []string `json:"errors"`
}

// NewAPIError returns a new instance of an APIError.
func NewAPIError(errors ...error) APIError {
	return &apiError{
		errToString(errors),
	}
}

func errToString(errors []error) []string {
	strs := []string{}

	for _, err := range errors {
		strs = append(strs, err.Error())
	}

	return strs
}
