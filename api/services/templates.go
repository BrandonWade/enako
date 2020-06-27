package services

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// TemplateService an interface for working with categories.
//go:generate counterfeiter -o fakes/fake_template_service.go . TemplateService
type TemplateService interface {
	GenerateAccountActivationEmail(link string) (string, error)
}

type templateService struct {
	logger *logrus.Logger
}

// NewTemplateService returns a new instance of a TemplateService.
func NewTemplateService(logger *logrus.Logger) TemplateService {
	return &templateService{
		logger,
	}
}

// GenerateAccountActivationEmail generates an account activation email with the provided link.
func (t *templateService) GenerateAccountActivationEmail(link string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		t.logger.WithFields(logrus.Fields{
			"method": "TemplateService.GenerateAccountActivationEmail",
			"link":   link,
		}).Error(err.Error())
		return "", err
	}

	path := filepath.Join(wd, "./templates/activate.tmpl")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		t.logger.WithFields(logrus.Fields{
			"method": "TemplateService.GenerateAccountActivationEmail",
			"link":   link,
			"path":   path,
		}).Error(err.Error())
		return "", err
	}

	data := struct {
		ActivationLink string
	}{
		link,
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		t.logger.WithFields(logrus.Fields{
			"method": "TemplateService.GenerateAccountActivationEmail",
			"link":   link,
			"path":   path,
		}).Error(err.Error())
		return "", err
	}

	return tpl.String(), nil
}
