package documents_validate

//
// helpers => validations => document => cnpj_test.go
//

import (
	"testing"
)

var (
	cnpjTests = []struct {
		cnpj        string
		cnpjIsValid bool
		onlyDigits  bool
	}{
		{cnpj: "3467875434578764345789654", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "", cnpjIsValid: false, onlyDigits: false},
		{cnpj: " ", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "AAAAAAAVBDDD", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "00.000.000/0000-00", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "11.111.111/1111-11", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "22.222.222/2222-22", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "33.333.333/3333-33", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "44.444.444/4444-44", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "55.555.555/5555-55", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "66.666.666/6666-66", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "77.777.777/7777-77", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "88.888.888/8888-88", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "99.999.999/9999-99", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "26.637.142/0001-85", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "74.221.325/0001-03", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "26-637-142.0001/58", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "74-221-325.0001/30", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "26.637.142/0001-58", cnpjIsValid: true, onlyDigits: false},
		{cnpj: "74.221.325/0001-30", cnpjIsValid: true, onlyDigits: false},
		{cnpj: "03.914.634/0001-48", cnpjIsValid: true, onlyDigits: false},
		{cnpj: "03914634000148", cnpjIsValid: true, onlyDigits: false},
		{cnpj: "03.914.634/0001-4", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "0391463400014", cnpjIsValid: false, onlyDigits: false},
		{cnpj: "03914634000148", cnpjIsValid: true, onlyDigits: true},
		{cnpj: "03.914.634/0001-48", cnpjIsValid: false, onlyDigits: true},
	}
)

func TestIsValidCNPJ(t *testing.T) {

	for _, c := range cnpjTests {
		isvalid := IsValidCNPJ(c.cnpj, c.onlyDigits)

		if isvalid && !c.cnpjIsValid {
			t.Errorf(`Erro o CNPJ "%s" INVALIDO foi considerado VALIDO`, c.cnpj)
		}

		if !isvalid && c.cnpjIsValid {
			t.Errorf(`Erro o CNPJ "%s" VALIDO foi considerado INVALIDO`, c.cnpj)
		}
	}
}
