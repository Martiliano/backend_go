package gomailv2

//
// helpers => email => gomailv2 => set_imap_server.go
//

func (e *Email) SetImapServer(server string, secure bool, port int) error {
	e.ImapPort = port
	e.ImapServer = server
	e.ImapPort = port

	return nil
}
