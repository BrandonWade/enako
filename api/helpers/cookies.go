package helpers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionCookieName the name of the browser session cookie
const SessionCookieName = "enako-session"

// CookieStorer an interface for working with a CookieStore
//go:generate counterfeiter -o fakes/fake_cookie_helper.go . CookieStorer
type CookieStorer interface {
	Get(r *http.Request, name string) (SessionStorer, error)
	IsAuthenticated(r *http.Request) (bool, error)
}

// CookieStore a wrapper around the gorilla/sessions CookieStore to allow mocking
type CookieStore struct {
	store *sessions.CookieStore
}

// NewCookieStore the constructor for a new CookieStore
func NewCookieStore(keys []byte) *CookieStore {
	return &CookieStore{
		sessions.NewCookieStore(keys),
	}
}

// Get retrieve the cookie with the given name from the underlying CookieStore
func (c *CookieStore) Get(r *http.Request, name string) (SessionStorer, error) {
	s, err := c.store.Get(r, name)
	if err != nil {
		return nil, err
	}

	return &SessionStore{s}, nil
}

// IsAuthenticated returns
func (c *CookieStore) IsAuthenticated(r *http.Request) (bool, error) {
	s, err := c.store.Get(r, SessionCookieName)
	if err != nil {
		return false, err
	}

	return s.Values["authenticated"] == true, nil
}
