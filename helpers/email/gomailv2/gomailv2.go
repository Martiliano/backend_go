package gomailv2

//
// helpers => email => gomailv2 => gomailv2.go
//

import (
	email "BackEnd_Api/helpers/email"

	"gopkg.in/gomail.v2"
)

type Email struct {
	message *gomail.Message

	UserName     string
	UserPassWord string

	SmtpServer string
	SmtpSecure bool
	SmtpPort   int

	PopServer string
	PopSecure bool
	PopPort   int

	ImapServer string
	ImapSecure bool
	ImapPort   int
}

func NewEmailService() email.IEmail {
	return &Email{
		message: gomail.NewMessage(),

		UserName:     "",
		UserPassWord: "",

		SmtpServer: "",
		SmtpSecure: true,
		SmtpPort:   0,

		PopServer: "",
		PopSecure: true,
		PopPort:   0,

		ImapServer: "",
		ImapSecure: true,
		ImapPort:   0,
	}
}
