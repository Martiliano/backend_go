package phone_validate

import "testing"

//
// helpers => validations => phone => phone_test.go
//

var (
	phoneTests = []struct {
		phone                   string
		phoneIsValid            bool
		onlyDigits              bool
		fullNationalNumber      bool
		fullInternationalNumber bool
	}{
		{phone: "3467875434578764345789654", phoneIsValid: false, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "", phoneIsValid: false, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: " ", phoneIsValid: false, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "AAAAAAAAAAA", phoneIsValid: false, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "054988776655", phoneIsValid: false, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "54X988776655", phoneIsValid: false, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "(054)988776655)", phoneIsValid: false, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "(054) 988776655)", phoneIsValid: false, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},

		{phone: "5554988776655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "54988776655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "988776655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "+55 (54) 98877-6655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "(54) 98877-6655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "98877-6655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "(54) 98877-6655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "54 98877-6655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: false},

		{phone: "+55 (54) 98877-6655", phoneIsValid: false, onlyDigits: true, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "5554988776655", phoneIsValid: true, onlyDigits: true, fullNationalNumber: false, fullInternationalNumber: false},
		{phone: "5554988776655", phoneIsValid: true, onlyDigits: true, fullNationalNumber: false, fullInternationalNumber: true},
		{phone: "54988776655", phoneIsValid: true, onlyDigits: true, fullNationalNumber: true, fullInternationalNumber: false},
		{phone: "+55 (54) 98877-6655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: false, fullInternationalNumber: true},
		{phone: "(54) 98877-6655", phoneIsValid: true, onlyDigits: false, fullNationalNumber: true, fullInternationalNumber: false},
	}
)

func TestIsValidCPF(t *testing.T) {

	for _, c := range phoneTests {
		isvalid := IsValidePhone(c.phone, c.onlyDigits, c.fullNationalNumber, c.fullInternationalNumber)

		if isvalid && !c.phoneIsValid {
			t.Errorf(`Erro o TELEFONE "%s" INVALIDO foi considerado VALIDO`, c.phone)
		}

		if !isvalid && c.phoneIsValid {
			t.Errorf(`Erro o TELEFONE "%s" VALIDO foi considerado INVALIDO`, c.phone)
		}
	}
}
