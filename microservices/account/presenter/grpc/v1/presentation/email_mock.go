package account

//
// microservices => account => presenter => grpc => v1 => presentation => email_mock.go
//
//

import (
	email "BackEnd_Api/helpers/email"
)

type Email struct {
	messageBody *string
}

func NewEmailMockService(messageBody *string) email.IEmail {
	return &Email{
		messageBody: messageBody,
	}
}

func (e *Email) SendEmailSetBody(contentType string, content string) error {
	*e.messageBody = content
	return nil
}

func (e *Email) SendEmail() error {
	return nil
}

func (e *Email) SendEmailAttachFile(file string) error {
	return nil
}

func (e *Email) Initialize() error {
	return nil
}

func (e *Email) SetEmailCredentials(user string, password string) error {
	return nil
}

func (e *Email) SetSmtpServer(server string, secure bool, port int) error {
	return nil
}

func (e *Email) SetPopServer(server string, secure bool, port int) error {
	return nil
}

func (e *Email) SetImapServer(server string, secure bool, port int) error {
	return nil
}

func (e *Email) SendEmailFrom(address string, name string) error {
	return nil
}

func (e *Email) SendEmailTo(address string, name string) error {
	return nil
}

func (e *Email) SendEmailCc(address string, name string) error {
	return nil
}

func (e *Email) SendEmailBcc(address string, name string) error {
	return nil
}

func (e *Email) SendEmailAddHeader(key string, value ...string) error {
	return nil
}

func (e *Email) Finally() error {
	return nil
}
