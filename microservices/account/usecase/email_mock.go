package account

//
// microservices => account => usecase => email_mock.go
//

import (
	email "BackEnd_Api/helpers/email"
	"fmt"
)

type Email struct {
	message     chan<- string
	messageBody string
}

func NewEmailMockService(message chan<- string) email.IEmail {
	fmt.Println("usecase NewEmailMockService 1 ")
	return &Email{
		message: message,
	}
}

func (e *Email) SendEmailSetBody(contentType string, content string) error {
	e.messageBody = content
	return nil
}

func (e *Email) SendEmail() error {
	e.message <- e.messageBody
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
