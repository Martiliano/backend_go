package gomailv2

//
// helpers => email => gomailv2 => set_email_to.go
//

func (e *Email) SendEmailBcc(address string, name string) error {
	e.message.SetAddressHeader("Bcc", address, name)

	return nil
}
