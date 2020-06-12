package middleware

import (
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/sirupsen/logrus"
)

type key int

const (
	// ContextUserAccountKey ...
	ContextUserAccountKey key = iota

	// ContextExpenseKey ...
	ContextExpenseKey

	// ContextUserAccountIDKey ...
	ContextUserAccountIDKey
)

// Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// MiddlewareStacker ...
type MiddlewareStacker interface {
	Apply(http.HandlerFunc, []Middleware) http.HandlerFunc
}

// MiddlewareStack ...
type MiddlewareStack struct {
	logger *logrus.Logger
	store  helpers.CookieStorer
}

// NewMiddlewareStack ...
func NewMiddlewareStack(logger *logrus.Logger, store helpers.CookieStorer) *MiddlewareStack {
	return &MiddlewareStack{
		logger,
		store,
	}
}

// Apply ...
func (s *MiddlewareStack) Apply(f http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}
