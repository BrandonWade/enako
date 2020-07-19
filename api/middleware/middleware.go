package middleware

import (
	"net/http"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/sirupsen/logrus"
)

type key int

const (
	// ContextCreateAccountKey a key used to store & retrieve a create account model from the request context.
	ContextCreateAccountKey key = iota

	// ContextExpenseKey a key used to store & retrieve an expense model from the request context.
	ContextExpenseKey

	// ContextAccountIDKey a key used to store & retrieve an account id from the request context.
	ContextAccountIDKey

	// ContextRequestPasswordResetKey a key used to store & retrieve a request password reset model from the request context.
	ContextRequestPasswordResetKey

	// ContextPasswordResetKey a key used to store & retrieve a password reset model from the request context.
	ContextPasswordResetKey
)

// Middleware is a type alias for working with middleware.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Stack used to inject dependencies into middleware.
type Stack struct {
	logger *logrus.Logger
	store  helpers.CookieStorer
}

// NewMiddlewareStack returns a new instance of a middleware stack.
func NewMiddlewareStack(logger *logrus.Logger, store helpers.CookieStorer) *Stack {
	return &Stack{
		logger,
		store,
	}
}

// Apply applies a middleware stack to the given HTTP handler.
func Apply(f http.HandlerFunc, middlewares []Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}
