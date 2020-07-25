package services_test

import (
	"errors"
	"io/ioutil"

	helperfakes "github.com/BrandonWade/enako/api/helpers/fakes"
	"github.com/BrandonWade/enako/api/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

var _ = Describe("TemplateService", func() {
	var (
		logger          *logrus.Logger
		templater       *helperfakes.FakeTemplater
		templateService services.TemplateService
		link            string
	)

	BeforeEach(func() {
		logger = logrus.New()
		logger.Out = ioutil.Discard

		templater = &helperfakes.FakeTemplater{}
		templateService = services.NewTemplateService(logger, templater, "/test/path")

		link = "foo.bar?t=123abc"
	})

	Describe("GenerateAccountActivationEmail", func() {
		Context("when generating an account activation email", func() {
			It("returns an error when one occurs while creating the template", func() {
				templater.GenerateTemplateReturns("", errors.New("templater error"))

				template, err := templateService.GenerateAccountActivationEmail(link)
				Expect(template).To(BeEmpty())
				Expect(err).To(HaveOccurred())
			})

			It("returns the template and no error", func() {
				templater.GenerateTemplateReturns("<html><body>Template</body></html>", nil)

				template, err := templateService.GenerateAccountActivationEmail(link)
				Expect(template).NotTo(BeEmpty())
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("GeneratePasswordResetEmail", func() {
		Context("when generating a password reset email", func() {
			It("returns an error when one occurs while creating the template", func() {
				templater.GenerateTemplateReturns("", errors.New("templater error"))

				template, err := templateService.GeneratePasswordResetEmail(link)
				Expect(template).To(BeEmpty())
				Expect(err).To(HaveOccurred())
			})

			It("returns the template and no error", func() {
				templater.GenerateTemplateReturns("<html><body>Template</body></html>", nil)

				template, err := templateService.GeneratePasswordResetEmail(link)
				Expect(template).NotTo(BeEmpty())
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("GeneratePasswordUpdatedEmail", func() {
		Context("when generating a password updated email", func() {
			It("returns an error when one occurs while creating the template", func() {
				templater.GenerateTemplateReturns("", errors.New("templater error"))

				template, err := templateService.GeneratePasswordUpdatedEmail(link)
				Expect(template).To(BeEmpty())
				Expect(err).To(HaveOccurred())
			})

			It("returns the template and no error", func() {
				templater.GenerateTemplateReturns("<html><body>Template</body></html>", nil)

				template, err := templateService.GeneratePasswordUpdatedEmail(link)
				Expect(template).NotTo(BeEmpty())
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
