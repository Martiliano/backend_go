package roles_validate

//
// helpers => validations => roles => roles.go
//

func IsValidRole(roles ...string) bool {
	roles_values := [...]string{"SemRestricoes", "Incluir", "Alterar", "Excluir", "ConsultarPublico", "ConsultarPrivado", "Consultar", "Visualizar"}

	isValid := true

	for _, r := range roles {

		naotem := true
		for _, rv := range roles_values {
			if r == rv {
				naotem = false
				break
			}
		}

		if naotem {
			isValid = false
		}

	}

	return isValid
}
