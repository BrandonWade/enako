package helpers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionStorer an interface for working with a Session
//go:generate counterfeiter -o fakes/fake_session_helper.go . SessionStorer
type SessionStorer interface {
	Get(key interface{}) interface{}
	Set(key, value interface{})
	Delete()
	Save(r *http.Request, w http.ResponseWriter) error
}

// Session a wrapper around the gorilla/sessions Session to allow mocking
type Session struct {
	session *sessions.Session
}

// Get a value from the Session
func (s *Session) Get(key interface{}) interface{} {
	return s.session.Values[key]
}

// Set a value to the Session
func (s *Session) Set(key, value interface{}) {
	s.session.Values[key] = value
}

// Delete the session
func (s *Session) Delete() {
	s.session.Options.MaxAge = -1
}

// Save any changes made to the Session
func (s *Session) Save(r *http.Request, w http.ResponseWriter) {
	s.session.Save(r, w)
}
