package gomailv2

//
// helpers => email => gomailv2 => set_email_add_header.go
//

func (e *Email) SendEmailAddHeader(key string, value ...string) error {
	e.message.SetHeader(key, value...)

	return nil
}
