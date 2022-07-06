package gomailv2

//
// helpers => email => gomailv2 => set_email_credentials.go
//

func (e *Email) SetEmailCredentials(user string, password string) error {

	e.UserName = user
	e.UserPassWord = password

	return nil
}
