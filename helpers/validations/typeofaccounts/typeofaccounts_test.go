package typeofaccounts_validate

//
// helpers => validations => typeofaccounts => typeofaccounts_test.go
//

import "testing"

var (
	typeofaccountsTests = []struct {
		typeofaccount        []string
		typeofaccountIsValid bool
	}{
		{typeofaccount: []string{"3467875434578764345789654"}, typeofaccountIsValid: false},
		{typeofaccount: []string{""}, typeofaccountIsValid: false},
		{typeofaccount: []string{" "}, typeofaccountIsValid: false},
		{typeofaccount: []string{"AAAAAAAAAAA"}, typeofaccountIsValid: false},
		{typeofaccount: []string{"054988776655"}, typeofaccountIsValid: false},
		{typeofaccount: []string{"054988776655", "Incluir", "Excluir", "AAAAAAAAAAA"}, typeofaccountIsValid: false},

		{typeofaccount: []string{"Desenvolvedor"}, typeofaccountIsValid: true},
		{typeofaccount: []string{"Administrador"}, typeofaccountIsValid: true},
		{typeofaccount: []string{"Supervisor"}, typeofaccountIsValid: true},
		{typeofaccount: []string{"Gerente"}, typeofaccountIsValid: true},
		{typeofaccount: []string{"Colaborador"}, typeofaccountIsValid: true},
		{typeofaccount: []string{"Usuario"}, typeofaccountIsValid: true},
		{typeofaccount: []string{"Visitante"}, typeofaccountIsValid: true},
		{typeofaccount: []string{"Desenvolvedor", "Administrador", "Supervisor", "Gerente"}, typeofaccountIsValid: true},
	}
)

func TestIsValidTypeOfAccount(t *testing.T) {

	for _, c := range typeofaccountsTests {
		isvalid := IsValidTypeOfAccount(c.typeofaccount...)

		if isvalid && !c.typeofaccountIsValid {
			t.Errorf(`Erro o TYPEOFACCOUNT "%s" INVALIDO foi considerado VALIDO`, c.typeofaccount)
		}

		if !isvalid && c.typeofaccountIsValid {
			t.Errorf(`Erro o TYPEOFACCOUNT "%s" VALIDO foi considerado INVALIDO`, c.typeofaccount)
		}
	}
}
