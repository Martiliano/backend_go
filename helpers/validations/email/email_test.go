package email_validate

//
// helpers => validations => email => email_test.go
//

import (
	"testing"
)

var (
	emailTests = []struct {
		emailAddress   string
		formatIsValid  bool
		accountIsValid bool
	}{
		{emailAddress: "florian@carrere.cc", formatIsValid: true, accountIsValid: true},
		{emailAddress: "support@g2mail.com", formatIsValid: true, accountIsValid: false},
		{emailAddress: " florian@carrere.cc", formatIsValid: false, accountIsValid: false},
		{emailAddress: "florian@carrere.cc ", formatIsValid: false, accountIsValid: false},
		{emailAddress: "test@912-wrong-domain902.com", formatIsValid: true, accountIsValid: false},
		{emailAddress: "0932910-qsdcqozuioqkdmqpeidj8793@gmail.com", formatIsValid: true, accountIsValid: false},
		{emailAddress: "@gmail.com", formatIsValid: false, accountIsValid: false},
		{emailAddress: "test@gmail@gmail.com", formatIsValid: false, accountIsValid: false},
		{emailAddress: "test test@gmail.com", formatIsValid: false, accountIsValid: false},
		{emailAddress: " test@gmail.com", formatIsValid: false, accountIsValid: false},
		{emailAddress: "test@wrong domain.com", formatIsValid: false, accountIsValid: false},
		{emailAddress: "é&ààà@gmail.com", formatIsValid: false, accountIsValid: false},
		{emailAddress: "admin@notarealdomain12345.com", formatIsValid: true, accountIsValid: false},
		{emailAddress: "a@gmail.xyz", formatIsValid: true, accountIsValid: false},
		{emailAddress: "", formatIsValid: false, accountIsValid: false},
		{emailAddress: "email@email", formatIsValid: false, accountIsValid: false},
		{emailAddress: "not-a-valid-email", formatIsValid: false, accountIsValid: false},
	}
)

func TestValidateEmailFormat(t *testing.T) {
	for _, s := range emailTests {
		err := ValidateEmailFormat(s.emailAddress)

		if err != nil && s.formatIsValid == true {
			t.Errorf(`"%s" => Erro invalido: "%v"`, s.emailAddress, err)
		}
		if err == nil && s.formatIsValid == false {
			t.Errorf(`"%s" => Email INVALIDO considerado VALIDO. Erro não identificado`, s.emailAddress)
		}
	}
}

func TestValidateEmailHost(t *testing.T) {
	for _, s := range emailTests {
		if !s.formatIsValid {
			continue
		}

		err := ValidateEmailHost(s.emailAddress)
		if err != nil && s.accountIsValid == true {
			t.Errorf(`"%s" => Erro invalido: "%v"`, s.emailAddress, err)
		}
		if err == nil && s.accountIsValid == false {
			t.Errorf(`"%s" => Erro não identificado`, s.emailAddress)
		}
	}
}

func TestValidateEmailMX(t *testing.T) {
	for _, s := range emailTests {
		if !s.formatIsValid {
			continue
		}

		err := ValidateEmailMX(s.emailAddress)
		if err != nil && s.accountIsValid == true {
			t.Errorf(`"%s" => Erro invalido: "%v"`, s.emailAddress, err)
		}
		if err == nil && s.accountIsValid == false {
			t.Errorf(`"%s" => Erro não identificado`, s.emailAddress)
		}
	}
}

func TestValidateEmailHostAndUser(t *testing.T) {
	var (
		serverHostName    = "your.smtp.server"
		serverMailAddress = "your.smtp.server.email"
	)
	for _, s := range emailTests {
		if !s.formatIsValid {
			continue
		}

		err := ValidateEmailHostAndUser(serverHostName, serverMailAddress, s.emailAddress)
		if err != nil && s.accountIsValid == true {
			t.Errorf(`"%s" => Erro invalido: "%v"`, s.emailAddress, err)
		}
		if err == nil && s.accountIsValid == false {
			t.Errorf(`"%s" => Erro não identificado`, s.emailAddress)
		}
	}
}
