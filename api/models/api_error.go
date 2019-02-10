package models

type APIError interface{}

type apiError struct {
	Errors []string `json:"errors"`
}

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
