package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/BrandonWade/enako/api/middleware"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware", func() {
	var (
		mw        []middleware.Middleware
		decorator func(http.ResponseWriter, *http.Request)
		w         *httptest.ResponseRecorder
		r         *http.Request
	)

	BeforeEach(func() {
		w = httptest.NewRecorder()

		mw = []middleware.Middleware{
			func(f http.HandlerFunc) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("123"))
					f(w, r)
				}
			},
			func(f http.HandlerFunc) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("abc"))
					f(w, r)
				}
			},
		}
		decorator = func(w http.ResponseWriter, r *http.Request) {}
	})

	Describe("Middleware", func() {
		Context("when applying a stack of middleware to a handler", func() {
			It("applies the middleware to the http handler in order without causing errors", func() {
				r = httptest.NewRequest("POST", "/v1/accounts", strings.NewReader("{foo}"))

				handler := middleware.Apply(decorator, mw)
				handler(w, r)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(strings.TrimSpace(w.Body.String())).To(BeEquivalentTo("abc123"))
			})
		})
	})
})
