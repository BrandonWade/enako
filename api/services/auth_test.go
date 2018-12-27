package services_test

import (
	"errors"

	"github.com/BrandonWade/enako/api/repositories/fakes"
	"github.com/BrandonWade/enako/api/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthService", func() {

	var (
		authRepo    *fakes.FakeAuthRepository
		authService services.AuthService
	)

	BeforeEach(func() {
		authRepo = &fakes.FakeAuthRepository{}
		authService = services.NewAuthService(authRepo)
	})

	Describe("CreateAccount", func() {

		Context("when creating a new account", func() {

			var (
				accountID = int64(18742356)
				email     = "test@test.com"
				password  = "testpassword123"
			)

			It("returns an error when an error is encountered while creating the account", func() {
				authRepo.CreateAccountReturns(0, errors.New("repo error"))

				id, err := authService.CreateAccount(email, password)
				Expect(authRepo.CreateAccountCallCount()).To(Equal(1))
				Expect(id).To(Equal(int64(0)))
				Expect(err).To(HaveOccurred())
			})

			It("returns the id of the new account row and no error", func() {
				authRepo.CreateAccountReturns(accountID, nil)

				id, err := authService.CreateAccount(email, password)
				Expect(authRepo.CreateAccountCallCount()).To(Equal(1))
				Expect(id).To(Equal(accountID))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
