package gomailv2

//
// helpers => email => gomailv2 => set_email_to.go
//

func (e *Email) SendEmailTo(address string, name string) error {
	e.message.SetAddressHeader("To", address, name)

	return nil
}
