package gomailv2

//
// helpers => email => gomailv2 => initialize.go
//

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func (e *Email) Initialize() error {
	e.message = gomail.NewMessage()

	if e.message != nil {
		fmt.Println("Initialize Email message is not nil")
	} else {
		fmt.Println("Initialize Email message is nil")
	}

	return nil
}
