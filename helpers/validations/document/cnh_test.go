package documents_validate

import "testing"

//
// helpers => validations => document => cnh_test.go
//

var (
	cnhTests = []struct {
		cnh        string
		cnhIsValid bool
	}{
		{cnh: "3467875434578764345789654", cnhIsValid: false},
		{cnh: "", cnhIsValid: false},
		{cnh: " ", cnhIsValid: false},
		{cnh: "SCCSDDFDFDF", cnhIsValid: false},
		{cnh: "02102234243", cnhIsValid: false},
		{cnh: "02102234142", cnhIsValid: false},
		{cnh: "13798941353", cnhIsValid: false},
		{cnh: "00676003001", cnhIsValid: false},
		{cnh: "067600300-1", cnhIsValid: false},
		{cnh: "0067600300-1", cnhIsValid: false},
		{cnh: "81952476011", cnhIsValid: true},
		{cnh: "33798941353", cnhIsValid: true},
		{cnh: "87222700600", cnhIsValid: true},
		{cnh: "45991167705", cnhIsValid: true},
		{cnh: "19595699996", cnhIsValid: true},
		{cnh: "00067600300", cnhIsValid: true},
		{cnh: "01537890901", cnhIsValid: true},
		{cnh: "01537890902", cnhIsValid: false},
	}
)

func TestIsValidCNH(t *testing.T) {

	for _, c := range cnhTests {
		isvalid := IsValidCNH(c.cnh)

		if isvalid && !c.cnhIsValid {
			t.Errorf(`Erro o CNH "%s" INVALIDO foi considerado VALIDO`, c.cnh)
		}

		if !isvalid && c.cnhIsValid {
			t.Errorf(`Erro o CNH "%s" VALIDO foi considerado INVALIDO`, c.cnh)
		}
	}
}
