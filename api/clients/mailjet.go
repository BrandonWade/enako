package clients

import (
	"fmt"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/sirupsen/logrus"
)

// MailjetClient an interface for working with accounts and sessions.
//go:generate counterfeiter -o fakes/fake_mailjet_client.go . MailjetClient
type MailjetClient interface {
	Send(message mailjet.InfoMessagesV31) error
}

type mailjetClient struct {
	logger *logrus.Logger
	client *mailjet.Client
}

// NewMailjetClient returns a new instance of an MailjetClient.
func NewMailjetClient(logger *logrus.Logger, client *mailjet.Client) MailjetClient {
	return &mailjetClient{
		logger,
		client,
	}
}

// Send sends the provided message.
func (c *mailjetClient) Send(message mailjet.InfoMessagesV31) error {
	messages := mailjet.MessagesV31{
		Info: []mailjet.InfoMessagesV31{message},
	}

	res, err := c.client.SendMailV31(&messages)
	if err != nil {
		c.logger.WithFields(logrus.Fields{
			"method":    "MailjetClient.Send",
			"recipient": message.To,
			"emailID":   message.CustomID,
		}).Error(err.Error())
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}
