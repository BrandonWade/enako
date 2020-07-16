package models

// MessageResponse contains a list of API messages.
type MessageResponse struct {
	Messages []Message `json:"messages"`
}

// Message a struct for formatting API messages.
type Message struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// MessagesFromStrings converts a list of strings to messages.
func MessagesFromStrings(messages ...string) *MessageResponse {
	msgs := []Message{}

	for _, msg := range messages {
		msgs = append(msgs, Message{msg, "info"})
	}

	return &MessageResponse{
		msgs,
	}
}

// MessagesFromErrors converts a list of errors to messages.
func MessagesFromErrors(errors ...error) *MessageResponse {
	msgs := []Message{}

	for _, err := range errors {
		msgs = append(msgs, Message{err.Error(), "error"})
	}

	return &MessageResponse{
		msgs,
	}
}
