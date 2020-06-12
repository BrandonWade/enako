package helpers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionStorer an interface for working with a SessionStore
//go:generate counterfeiter -o fakes/fake_session_helper.go . SessionStorer
type SessionStorer interface {
	Get(key interface{}) interface{}
	Set(key, value interface{})
	Delete()
	Save(r *http.Request, w http.ResponseWriter)
}

// SessionStore a wrapper around the gorilla/sessions SessionStore to allow mocking
type SessionStore struct {
	session *sessions.Session
}

// Get a value from the SessionStore
func (s *SessionStore) Get(key interface{}) interface{} {
	return s.session.Values[key]
}

// Set a value to the SessionStore
func (s *SessionStore) Set(key, value interface{}) {
	s.session.Values[key] = value
}

// Delete the sessionStore
func (s *SessionStore) Delete() {
	s.session.Options.MaxAge = -1
}

// Save any changes made to the SessionStore
func (s *SessionStore) Save(r *http.Request, w http.ResponseWriter) {
	s.session.Save(r, w)
}
