package models

// APIMessage a wrapper interface for formatting messages.
type APIMessage interface{}

type apiMessage struct {
	Messages []string `json:"messages"`
}

// NewAPIMessage returns a new instance of an APIMessage.
func NewAPIMessage(messages ...string) APIMessage {
	return &apiMessage{
		messages,
	}
}
