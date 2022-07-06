package account

//
// microservices => account => controlled.go
//

import (
	auth "BackEnd_Api/microservices/auth/rules"
)

func GetAccountControlled() map[string]auth.Controlled {
	const microservicePath = "/iuris.account.v1.Account/"

	return map[string]auth.Controlled{
		microservicePath + "CreateAccount": {
			TypeOfAccount: []string{"Desenvolvedor", "CEO", "Presidente", "Conselheiro", "Administrador", "Gerente", "Supervisor", "Colaborador", "Usuario", "Visitante"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		microservicePath + "GetAccountById": {
			TypeOfAccount: []string{"Desenvolvedor", "CEO", "Presidente", "Conselheiro", "Administrador", "Gerente", "Supervisor", "Colaborador", "Usuario", "Visitante"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		microservicePath + "GetAccountByEmail": {
			TypeOfAccount: []string{"Desenvolvedor", "CEO", "Presidente", "Conselheiro", "Administrador", "Gerente", "Supervisor", "Colaborador", "Usuario", "Visitante"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		microservicePath + "GetAccountByPhone": {
			TypeOfAccount: []string{"Desenvolvedor", "CEO", "Presidente", "Conselheiro", "Administrador", "Gerente", "Supervisor", "Colaborador", "Usuario", "Visitante"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		microservicePath + "GetAccountByDocument": {
			TypeOfAccount: []string{"Desenvolvedor", "CEO", "Presidente", "Conselheiro", "Administrador", "Gerente", "Supervisor", "Colaborador", "Usuario", "Visitante"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		microservicePath + "GetAllAccounts": {
			TypeOfAccount: []string{"Desenvolvedor", "CEO", "Presidente", "Conselheiro", "Administrador", "Gerente", "Supervisor", "Colaborador", "Usuario", "Visitante"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		microservicePath + "UpdateAccount": {
			TypeOfAccount: []string{"Desenvolvedor", "CEO", "Presidente", "Conselheiro", "Administrador", "Gerente", "Supervisor", "Colaborador", "Usuario", "Visitante"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		microservicePath + "DeleteAccount": {
			TypeOfAccount: []string{"Desenvolvedor", "CEO", "Presidente", "Conselheiro", "Administrador", "Gerente", "Supervisor", "Colaborador", "Usuario", "Visitante"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	}

}
