package helpers

import (
	"github.com/dchest/uniuri"
)

// TokenGenerator an interface for generating tokens.
//go:generate counterfeiter -o fakes/fake_token_generator.go . TokenGenerator
type TokenGenerator interface {
	CreateToken(length int) string
}

type tokenGenerator struct{}

// NewTokenGenerator returns a new instance of a TokenGenerator.
func NewTokenGenerator() TokenGenerator {
	return &tokenGenerator{}
}

// CreateToken creates a token of the given length.
func (a *tokenGenerator) CreateToken(length int) string {
	return uniuri.NewLen(length)
}
