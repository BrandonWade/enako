package services_test

import (
	"errors"
	"io/ioutil"

	clientfakes "github.com/BrandonWade/enako/api/clients/fakes"
	"github.com/BrandonWade/enako/api/services"
	servicefakes "github.com/BrandonWade/enako/api/services/fakes"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EmailService", func() {
	var (
		logger          *logrus.Logger
		templateService *servicefakes.FakeTemplateService
		emailClient     *clientfakes.FakeMailjetClient
		emailService    services.EmailService
		email           string
		token           string
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		templateService = &servicefakes.FakeTemplateService{}
		emailClient = &clientfakes.FakeMailjetClient{}
		emailService = services.NewEmailService(logger, templateService, emailClient)

		email = "foo@bar.net"
		token = "thisisareallylongtokenthatneedstobesuperlongtopassvalidation1234"
	})

	Describe("SendAccountActivationEmail", func() {
		Context("when sending an account activation email", func() {
			It("returns an error if one occurred while generating the email", func() {
				templateService.GenerateAccountActivationEmailReturns("", errors.New("template error"))

				err := emailService.SendAccountActivationEmail(email, token)
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while sending the email", func() {
				templateService.GenerateAccountActivationEmailReturns("email template", nil)
				emailClient.SendReturns(errors.New("client error"))

				err := emailService.SendAccountActivationEmail(email, token)
				Expect(err).To(HaveOccurred())
			})

			It("sends the email and returns no error", func() {
				templateService.GenerateAccountActivationEmailReturns("email template", nil)
				emailClient.SendReturns(nil)

				err := emailService.SendAccountActivationEmail(email, token)
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("SendPasswordResetEmail", func() {
		Context("when sending a password reset email", func() {
			It("returns an error if one occurred while generating the email", func() {
				templateService.GeneratePasswordResetEmailReturns("", errors.New("template error"))

				err := emailService.SendPasswordResetEmail(email, token)
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while sending the email", func() {
				templateService.GeneratePasswordResetEmailReturns("email template", nil)
				emailClient.SendReturns(errors.New("client error"))

				err := emailService.SendPasswordResetEmail(email, token)
				Expect(err).To(HaveOccurred())
			})

			It("sends the email and returns no error", func() {
				templateService.GeneratePasswordResetEmailReturns("email template", nil)
				emailClient.SendReturns(nil)

				err := emailService.SendPasswordResetEmail(email, token)
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("SendPasswordUpdatedEmail", func() {
		Context("when sending a password updated email", func() {
			It("returns an error if one occurred while generating the email", func() {
				templateService.GeneratePasswordUpdatedEmailReturns("", errors.New("template error"))

				err := emailService.SendPasswordUpdatedEmail(email)
				Expect(err).To(HaveOccurred())
			})

			It("returns an error if one occurred while sending the email", func() {
				templateService.GeneratePasswordUpdatedEmailReturns("email template", nil)
				emailClient.SendReturns(errors.New("client error"))

				err := emailService.SendPasswordUpdatedEmail(email)
				Expect(err).To(HaveOccurred())
			})

			It("sends the email and returns no error", func() {
				templateService.GeneratePasswordUpdatedEmailReturns("email template", nil)
				emailClient.SendReturns(nil)

				err := emailService.SendPasswordUpdatedEmail(email)
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
