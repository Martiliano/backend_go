package gomailv2

//
// helpers => email => gomailv2 => set_email_body.go
//

func (e *Email) SendEmailSetBody(contentType string, content string) error {
	e.message.SetBody(contentType, content)

	return nil
}
