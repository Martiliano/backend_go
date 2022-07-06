package gomailv2

import "fmt"

//
// helpers => email => gomailv2 => set_email_from.go
//

func (e *Email) SendEmailFrom(address string, name string) error {

	if e.message != nil {
		fmt.Println("Email message is not nil")
		e.message.SetAddressHeader("From", address, name)
	} else {
		fmt.Println("Email message is nil")
	}

	return nil
}
