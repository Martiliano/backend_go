package documents_validate

import "testing"

//
// helpers => validations => document => cns_test.go
//

var (
	cnsTests = []struct {
		cns        string
		cnsIsValid bool
	}{
		{cns: "3467875434578764345789654", cnsIsValid: false},
		{cns: "", cnsIsValid: false},
		{cns: " ", cnsIsValid: false},
		{cns: "SCCSDDFDFDF", cnsIsValid: false},
		{cns: "915 5017 0193 0306", cnsIsValid: false},
		{cns: "174 2241 7133 0004", cnsIsValid: false},
		{cns: "259 7557 3388 0001", cnsIsValid: false},
		{cns: "808-2536-1743-0486", cnsIsValid: false},
		{cns: "9999 0236 0200 834", cnsIsValid: false},
		{cns: "174 5984 3528 0018", cnsIsValid: true},
		{cns: "259 9557 3388 0001", cnsIsValid: true},
		{cns: "750 6557 1844 0005", cnsIsValid: true},
		{cns: "750655718440005", cnsIsValid: true},
	}
)

func TestIsValidCNS(t *testing.T) {

	for _, c := range cnsTests {
		isvalid := IsValidCNS(c.cns)

		if isvalid && !c.cnsIsValid {
			t.Errorf(`Erro o CNS "%s" INVALIDO foi considerado VALIDO`, c.cns)
		}

		if !isvalid && c.cnsIsValid {
			t.Errorf(`Erro o CNS "%s" VALIDO foi considerado INVALIDO`, c.cns)
		}
	}
}
