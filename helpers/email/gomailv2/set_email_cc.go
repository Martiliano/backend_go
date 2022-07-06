package gomailv2

//
// helpers => email => gomailv2 => set_email_to.go
//

func (e *Email) SendEmailCc(address string, name string) error {
	e.message.SetAddressHeader("Cc", address, name)

	return nil
}
