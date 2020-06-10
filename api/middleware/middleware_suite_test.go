package middleware_test

import (
	"testing"

	"github.com/BrandonWade/enako/api/validation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Middleware Suite")
}

var _ = BeforeSuite(func() {
	validation.InitValidator()
})
