package gomailv2

//
// helpers => email => gomailv2 => send_email.go
//

import (
	"gopkg.in/gomail.v2"
)

func (e *Email) SendEmail() error {

	d := gomail.NewDialer(e.SmtpServer, e.SmtpPort, e.UserName, e.UserPassWord)

	if err := d.DialAndSend(e.message); err != nil {
		return err
	}

	return nil
}
