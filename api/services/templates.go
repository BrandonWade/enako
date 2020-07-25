package services

import (
	"path/filepath"

	"github.com/BrandonWade/enako/api/helpers"
	"github.com/sirupsen/logrus"
)

// TemplateService an interface for working with categories.
//go:generate counterfeiter -o fakes/fake_template_service.go . TemplateService
type TemplateService interface {
	GenerateAccountActivationEmail(link string) (string, error)
	GeneratePasswordResetEmail(link string) (string, error)
	GeneratePasswordUpdatedEmail(link string) (string, error)
}

type templateService struct {
	logger    *logrus.Logger
	templater helpers.Templater
	basePath  string
}

// NewTemplateService returns a new instance of a TemplateService.
func NewTemplateService(logger *logrus.Logger, templater helpers.Templater, basePath string) TemplateService {
	return &templateService{
		logger,
		templater,
		basePath,
	}
}

// GenerateAccountActivationEmail generates an account activation email with the provided link.
func (t *templateService) GenerateAccountActivationEmail(link string) (string, error) {
	path := filepath.Join(t.basePath, "./templates/activate.tmpl")
	data := struct {
		ActivationLink string
	}{
		link,
	}

	return t.templater.GenerateTemplate(path, data)
}

// GeneratePasswordResetEmail generates a password reset email with the provided link.
func (t *templateService) GeneratePasswordResetEmail(link string) (string, error) {
	path := filepath.Join(t.basePath, "./templates/reset_password.tmpl")
	data := struct {
		PasswordResetLink string
	}{
		link,
	}

	return t.templater.GenerateTemplate(path, data)
}

// GeneratePasswordUpdatedEmail generates an email when a password was updated.
func (t *templateService) GeneratePasswordUpdatedEmail(link string) (string, error) {
	path := filepath.Join(t.basePath, "./templates/password_updated.tmpl")
	data := struct {
		ForgotPasswordLink string
	}{
		link,
	}

	return t.templater.GenerateTemplate(path, data)
}
