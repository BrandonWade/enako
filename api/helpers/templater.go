package helpers

import (
	"bytes"
	"html/template"

	"github.com/sirupsen/logrus"
)

// Templater an interface for parsing and generating templates.
//go:generate counterfeiter -o fakes/fake_templater.go . Templater
type Templater interface {
	GenerateTemplate(path string, data interface{}) (string, error)
}

type templater struct {
	logger *logrus.Logger
}

// NewTemplater returns a new instance of a Templater.
func NewTemplater(logger *logrus.Logger) Templater {
	return &templater{
		logger,
	}
}

func (t *templater) GenerateTemplate(path string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		t.logger.WithFields(logrus.Fields{
			"method": "Templater.GenerateTemplate",
			"path":   path,
		}).Error(err.Error())
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		t.logger.WithFields(logrus.Fields{
			"method": "Templater.GenerateTemplate",
			"path":   path,
		}).Error(err.Error())
		return "", err
	}

	return buf.String(), nil
}
