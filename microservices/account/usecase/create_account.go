package account

//
// microservices => account => usecase => create_account.go
//

import (
	documents_validate "BackEnd_Api/helpers/validations/document"
	email_validate "BackEnd_Api/helpers/validations/email"
	phone_validate "BackEnd_Api/helpers/validations/phone"
	roles_validate "BackEnd_Api/helpers/validations/roles"
	typeofaccounts_validate "BackEnd_Api/helpers/validations/typeofaccounts"
	account "BackEnd_Api/microservices/account/rules"
	"errors"
)

func (aUC *AccountUseCase) CreateAccount(a account.Account) (string, error) {
	err := email_validate.ValidateEmailFormat(a.Email)

	if err != nil {
		return "", err
	}

	if !phone_validate.IsValidePhone(a.PhoneNumber, true, true, false) {
		return "", errors.New("Telefone Formato Invalido")
	}

	if len(a.Document) > 0 {
		if !documents_validate.IsValidCPF(a.Document, true) && !documents_validate.IsValidCNPJ(a.Document, true) {
			return "", errors.New("CPF ou CNPJ Invalido")
		}
	}

	if !roles_validate.IsValidRole(a.Roles...) {
		return "", errors.New("Roles/Permissões Invalidos")
	}

	if !typeofaccounts_validate.IsValidTypeOfAccount(a.TypeOfAccount...) {
		return "", errors.New("Tipo de Contas Invalidos")
	}

	_, err = aUC.GetAccountByEmail(a.Email)

	if err == nil {
		return "", errors.New("Erro já existe um usuário utilizando este e-mail: " + a.Email)
	}

	_, err = aUC.GetAccountByPhone(a.PhoneNumber)

	if err == nil {
		return "", errors.New("Erro já existe um usuário utilizando este telefone: " + a.PhoneNumber)
	}

	if len(a.Document) > 0 {
		_, err = aUC.GetAccountByDocument(a.Document)

		if err == nil {
			return "", errors.New("Erro já existe um usuário utilizando este documento: " + a.Document)
		}
	}

	code, err := aUC.repository.CreateAccountRepository(a)

	if err != nil {
		return "", err
	}

	return code, nil
}
