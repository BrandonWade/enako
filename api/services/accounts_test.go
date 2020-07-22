package services_test

import (
	"errors"
	"io/ioutil"

	helpers "github.com/BrandonWade/enako/api/helpers/fakes"
	"github.com/BrandonWade/enako/api/models"
	"github.com/BrandonWade/enako/api/repositories/fakes"
	"github.com/BrandonWade/enako/api/services"
	servicefakes "github.com/BrandonWade/enako/api/services/fakes"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AccountService", func() {
	var (
		logger         *logrus.Logger
		hasher         *helpers.FakePasswordHasher
		generator      *helpers.FakeTokenGenerator
		emailService   *servicefakes.FakeEmailService
		accountRepo    *fakes.FakeAccountRepository
		accountService services.AccountService
		accountID      int64
		token          string
		email          string
		password       string
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		hasher = &helpers.FakePasswordHasher{}
		generator = &helpers.FakeTokenGenerator{}
		emailService = &servicefakes.FakeEmailService{}
		accountRepo = &fakes.FakeAccountRepository{}
		accountService = services.NewAccountService(logger, hasher, generator, emailService, accountRepo)

		accountID = int64(18742356)
		token = "thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234"
		email = "foo@bar.net"
		password = "testpassword123"
	})

	Describe("CreateAccount", func() {
		Context("when creating a new account", func() {
			It("returns an error when one occurs while hashing a password", func() {
				hasher.GenerateReturns("", errors.New("hasher error"))

				id, err := accountService.CreateAccount(email, password)
				Expect(hasher.GenerateCallCount()).To(Equal(1))
				Expect(id).To(BeEquivalentTo(0))
				Expect(err).To(HaveOccurred())
			})

			It("returns an error when the repo returns an error", func() {
				hasher.GenerateReturns("hashedtestpassword", nil)
				accountRepo.CreateAccountReturns(0, errors.New("repo error"))

				id, err := accountService.CreateAccount(email, password)
				Expect(accountRepo.CreateAccountCallCount()).To(Equal(1))
				Expect(id).To(BeEquivalentTo(0))
				Expect(err).To(HaveOccurred())
			})

			It("returns the id of the new account and no error", func() {
				hasher.GenerateReturns("hashedtestpassword", nil)
				accountRepo.CreateAccountReturns(accountID, nil)

				id, err := accountService.CreateAccount(email, password)
				Expect(accountRepo.CreateAccountCallCount()).To(Equal(1))
				Expect(id).To(Equal(accountID))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("RegisterUser", func() {
		Context("when registering a new user", func() {
			It("returns an error when one occurs while creating an account", func() {
				hasher.GenerateReturns("hashedtestpassword", nil)
				accountRepo.CreateAccountReturns(0, errors.New("repo error"))

				id, err := accountService.RegisterUser(email, password)
				Expect(id).To(BeEquivalentTo(0))
				Expect(err).To(HaveOccurred())
			})

			It("returns an error when one occurs while creating an activation token", func() {
				hasher.GenerateReturns("hashedtestpassword", nil)
				accountRepo.CreateAccountReturns(accountID, nil)
				generator.CreateTokenReturns(token)
				accountRepo.CreateActivationTokenReturns(0, errors.New("repo error"))

				id, err := accountService.RegisterUser(email, password)
				Expect(id).To(BeEquivalentTo(0))
				Expect(err).To(HaveOccurred())
			})

			It("returns an error when one occurs while sending an account activation email", func() {
				hasher.GenerateReturns("hashedtestpassword", nil)
				accountRepo.CreateAccountReturns(accountID, nil)
				generator.CreateTokenReturns(token)
				accountRepo.CreateActivationTokenReturns(1, nil)
				emailService.SendAccountActivationEmailReturns(errors.New("email error"))

				id, err := accountService.RegisterUser(email, password)
				Expect(id).To(BeEquivalentTo(0))
				Expect(err).To(HaveOccurred())
			})

			It("returns the id of the new account and no error", func() {
				hasher.GenerateReturns("hashedtestpassword", nil)
				accountRepo.CreateAccountReturns(accountID, nil)
				generator.CreateTokenReturns(token)
				accountRepo.CreateActivationTokenReturns(1, nil)
				emailService.SendAccountActivationEmailReturns(nil)

				id, err := accountService.RegisterUser(email, password)
				Expect(id).To(BeEquivalentTo(accountID))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("VerifyAccount", func() {
		Context("when verifying an activated account with the given credentials exists", func() {
			It("returns an error if one occurred while retrieving the account from the repo", func() {
				accountRepo.GetAccountByEmailReturns(&models.Account{}, errors.New("repo error"))

				id, err := accountService.VerifyAccount(email, password)
				Expect(id).To(BeEquivalentTo(0))
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while checking the given password against an account", func() {
				account := &models.Account{ID: accountID, Email: email, Password: password}
				accountRepo.GetAccountByEmailReturns(account, nil)
				hasher.CompareReturns(errors.New("hasher error"))

				id, err := accountService.VerifyAccount(email, password)
				Expect(id).To(BeEquivalentTo(0))
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while retrieving the activation token for the account", func() {
				account := &models.Account{ID: accountID, Email: email, Password: password, IsActivated: false}
				accountRepo.GetAccountByEmailReturns(account, nil)
				hasher.CompareReturns(nil)
				accountRepo.GetActivationTokenByAccountIDReturns(&models.ActivationToken{}, errors.New("repo error"))

				id, err := accountService.VerifyAccount(email, password)
				Expect(id).To(BeEquivalentTo(0))
				Expect(err).To(HaveOccurred())
			})

			It("returns the id of the new account and no error", func() {
				account := &models.Account{ID: accountID, Email: email, Password: password, IsActivated: true}
				accountRepo.GetAccountByEmailReturns(account, nil)
				hasher.CompareReturns(nil)

				id, err := accountService.VerifyAccount(email, password)
				Expect(id).To(BeEquivalentTo(accountID))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
