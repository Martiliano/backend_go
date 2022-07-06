package roles_validate

//
// helpers => validations => roles => roles_test.go
//

import "testing"

var (
	rolesTests = []struct {
		role        []string
		roleIsValid bool
	}{
		{role: []string{"3467875434578764345789654"}, roleIsValid: false},
		{role: []string{""}, roleIsValid: false},
		{role: []string{" "}, roleIsValid: false},
		{role: []string{"AAAAAAAAAAA"}, roleIsValid: false},
		{role: []string{"054988776655"}, roleIsValid: false},
		{role: []string{"054988776655", "Incluir", "Excluir", "AAAAAAAAAAA"}, roleIsValid: false},

		{role: []string{"SemRestricoes"}, roleIsValid: true},
		{role: []string{"Incluir"}, roleIsValid: true},
		{role: []string{"Alterar"}, roleIsValid: true},
		{role: []string{"Excluir"}, roleIsValid: true},
		{role: []string{"ConsultarPublico"}, roleIsValid: true},
		{role: []string{"ConsultarPrivado"}, roleIsValid: true},
		{role: []string{"Visualizar"}, roleIsValid: true},
		{role: []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"}, roleIsValid: true},
	}
)

func TestIsValidRole(t *testing.T) {

	for _, c := range rolesTests {
		isvalid := IsValidRole(c.role...)

		if isvalid && !c.roleIsValid {
			t.Errorf(`Erro o ROLE "%s" INVALIDO foi considerado VALIDO`, c.role)
		}

		if !isvalid && c.roleIsValid {
			t.Errorf(`Erro o ROLE "%s" VALIDO foi considerado INVALIDO`, c.role)
		}
	}
}
