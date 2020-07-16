package models

// Message a struct for formatting API messages.
type Message struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// MessagesFromStrings converts a list of strings to messages.
func MessagesFromStrings(messages ...string) []Message {
	msgs := []Message{}

	for _, msg := range messages {
		msgs = append(msgs, Message{msg, "info"})
	}

	return msgs
}

// MessagesFromErrors converts a list of errors to messages.
func MessagesFromErrors(errors ...error) []Message {
	msgs := []Message{}

	for _, err := range errors {
		msgs = append(msgs, Message{err.Error(), "error"})
	}

	return msgs
}
