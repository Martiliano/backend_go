package typeofaccounts_validate

//
// helpers => validations => roles => typeofaccounts.go
//

func IsValidTypeOfAccount(accounts ...string) bool {
	account_values := [...]string{"Desenvolvedor", "Administrador", "Supervisor", "Gerente", "Colaborador", "Usuario", "Visitante"}

	isValid := true

	for _, r := range accounts {

		naotem := true
		for _, rv := range account_values {
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
