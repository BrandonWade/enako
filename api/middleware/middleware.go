package middleware

import (
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/sirupsen/logrus"
)

type key int

const (
	// ContextUserAccountKey a key used to store & retrieve a user account from the request context.
	ContextUserAccountKey key = iota

	// ContextExpenseKey a key used to store & retrieve an expense from the request context.
	ContextExpenseKey

	// ContextUserAccountIDKey a key used to store & retrieve an account id from the request context.
	ContextUserAccountIDKey
)

// Middleware is a type alias for working with middleware.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// MiddlewareStacker an interface for working with a middleware stack.
type MiddlewareStacker interface {
	Apply(http.HandlerFunc, []Middleware) http.HandlerFunc
}

// MiddlewareStack used to inject dependencies into middleware.
type MiddlewareStack struct {
	logger *logrus.Logger
	store  helpers.CookieStorer
}

// NewMiddlewareStack returns a new instance of a MiddlewareStack.
func NewMiddlewareStack(logger *logrus.Logger, store helpers.CookieStorer) *MiddlewareStack {
	return &MiddlewareStack{
		logger,
		store,
	}
}

// Apply applies a middleware stack to the given HTTP handler.
func (s *MiddlewareStack) Apply(f http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}
