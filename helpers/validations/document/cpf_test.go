package documents_validate

//
// helpers => validations => document => cpf_test.go
//

import (
	"testing"
)

var (
	cpfTests = []struct {
		cpf        string
		cpfIsValid bool
		onlyDigits bool
	}{
		{cpf: "3467875434578764345789654", cpfIsValid: false, onlyDigits: false},
		{cpf: "", cpfIsValid: false, onlyDigits: false},
		{cpf: " ", cpfIsValid: false, onlyDigits: false},
		{cpf: "AAAAAAAAAAA", cpfIsValid: false, onlyDigits: false},
		{cpf: "000.000.000-00", cpfIsValid: false, onlyDigits: false},
		{cpf: "111.111.111-11", cpfIsValid: false, onlyDigits: false},
		{cpf: "222.222.222-22", cpfIsValid: false, onlyDigits: false},
		{cpf: "333.333.333-33", cpfIsValid: false, onlyDigits: false},
		{cpf: "444.444.444-44", cpfIsValid: false, onlyDigits: false},
		{cpf: "555.555.555-55", cpfIsValid: false, onlyDigits: false},
		{cpf: "666.666.666-66", cpfIsValid: false, onlyDigits: false},
		{cpf: "777.777.777-77", cpfIsValid: false, onlyDigits: false},
		{cpf: "888.888.888-88", cpfIsValid: false, onlyDigits: false},
		{cpf: "999.999.999-99", cpfIsValid: false, onlyDigits: false},
		{cpf: "248.438.034-08", cpfIsValid: false, onlyDigits: false},
		{cpf: "099.075.865-06", cpfIsValid: false, onlyDigits: false},
		{cpf: "248 438 034 8", cpfIsValid: false, onlyDigits: false},
		{cpf: "099-075-865.60", cpfIsValid: false, onlyDigits: false},
		{cpf: "099-075-865.6", cpfIsValid: false, onlyDigits: false},
		{cpf: "6384921600", cpfIsValid: false, onlyDigits: false},
		{cpf: "248.438.034-80", cpfIsValid: true, onlyDigits: false},
		{cpf: "099.075.865-60", cpfIsValid: true, onlyDigits: false},
		{cpf: "63849216004", cpfIsValid: true, onlyDigits: false},
		{cpf: "63849216004", cpfIsValid: true, onlyDigits: true},
		{cpf: "638.492.160-04", cpfIsValid: false, onlyDigits: true},
	}
)

func TestIsValidCPF(t *testing.T) {

	for _, c := range cpfTests {
		isvalid := IsValidCPF(c.cpf, c.onlyDigits)

		if isvalid && !c.cpfIsValid {
			t.Errorf(`Erro o CPF "%s" INVALIDO foi considerado VALIDO`, c.cpf)
		}

		if !isvalid && c.cpfIsValid {
			t.Errorf(`Erro o CPF "%s" VALIDO foi considerado INVALIDO`, c.cpf)
		}
	}
}
