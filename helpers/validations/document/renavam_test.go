package documents_validate

import "testing"

//
// helpers => validations => document => renavam_test.go
//

var (
	renavamTests = []struct {
		renavam        string
		renavamIsValid bool
	}{
		{renavam: "3467875434578764345789654", renavamIsValid: false},
		{renavam: "", renavamIsValid: false},
		{renavam: " ", renavamIsValid: false},
		{renavam: "SCCSDDFDFDF", renavamIsValid: false},
		{renavam: "38872054170", renavamIsValid: false},
		{renavam: "40999838209", renavamIsValid: false},
		{renavam: "31789431480", renavamIsValid: false},
		{renavam: "38919643060", renavamIsValid: false},
		{renavam: "38918883060", renavamIsValid: false},
		{renavam: "13824652268", renavamIsValid: true},
		{renavam: "08543317523", renavamIsValid: true},
		{renavam: "09769017014", renavamIsValid: true},
		{renavam: "81952476011", renavamIsValid: false},
		{renavam: "01993520012", renavamIsValid: true},
		{renavam: "04598137389", renavamIsValid: true},
		{renavam: "05204907510", renavamIsValid: true},
	}
)

func TestIsValidRENAVAM(t *testing.T) {

	for _, c := range renavamTests {
		isvalid := IsValidRENAVAM(c.renavam)

		if isvalid && !c.renavamIsValid {
			t.Errorf(`Erro o RENAVAM "%s" INVALIDO foi considerado VALIDO`, c.renavam)
		}

		if !isvalid && c.renavamIsValid {
			t.Errorf(`Erro o RENAVAM "%s" VALIDO foi considerado INVALIDO`, c.renavam)
		}
	}
}
