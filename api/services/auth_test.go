package services_test

import (
	"errors"
	"io/ioutil"

	helpers "github.com/BrandonWade/enako/api/helpers/fakes"
	"github.com/BrandonWade/enako/api/repositories/fakes"
	"github.com/BrandonWade/enako/api/services"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthService", func() {
	var (
		logger      *logrus.Logger
		hasher      *helpers.FakePasswordHasher
		authRepo    *fakes.FakeAuthRepository
		authService services.AuthService
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		hasher = &helpers.FakePasswordHasher{}
		authRepo = &fakes.FakeAuthRepository{}
		authService = services.NewAuthService(logger, hasher, authRepo)
	})

	Describe("CreateAccount", func() {
		Context("when creating a new account", func() {
			var (
				accountID = int64(18742356)
				username  = "foobar"
				email     = "test@test.com"
				password  = "testpassword123"
			)

			It("returns an error when a hasher error is encountered", func() {
				hasher.GenerateReturns("", errors.New("hasher error"))

				id, err := authService.CreateAccount(username, email, password)
				Expect(hasher.GenerateCallCount()).To(Equal(1))
				Expect(id).To(Equal(int64(0)))
				Expect(err).To(HaveOccurred())
			})

			It("returns an error when a repo error is encountered", func() {
				hasher.GenerateReturns("hashedtestpassword", nil)
				authRepo.CreateAccountReturns(0, errors.New("repo error"))

				id, err := authService.CreateAccount(username, email, password)
				Expect(authRepo.CreateAccountCallCount()).To(Equal(1))
				Expect(id).To(Equal(int64(0)))
				Expect(err).To(HaveOccurred())
			})

			It("returns the id of the new account row and no error", func() {
				hasher.GenerateReturns("hashedtestpassword", nil)
				authRepo.CreateAccountReturns(accountID, nil)

				id, err := authService.CreateAccount(username, email, password)
				Expect(authRepo.CreateAccountCallCount()).To(Equal(1))
				Expect(id).To(Equal(accountID))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
