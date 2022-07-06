package gomailv2

//
// helpers => email => gomailv2 => set_pop_server.go
//

func (e *Email) SetPopServer(server string, secure bool, port int) error {
	e.PopPort = port
	e.PopServer = server
	e.PopPort = port

	return nil
}
