package gomailv2

//
// helpers => email => gomailv2 => set_smtp_server.go
//

func (e *Email) SetSmtpServer(server string, secure bool, port int) error {
	e.SmtpPort = port
	e.SmtpServer = server
	e.SmtpPort = port

	return nil
}
