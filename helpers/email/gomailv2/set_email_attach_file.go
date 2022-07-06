package gomailv2

//
// helpers => email => gomailv2 => set_email_attach_file.go
//

func (e *Email) SendEmailAttachFile(file string) error {
	e.message.Attach(file)

	return nil
}
