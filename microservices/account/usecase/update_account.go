package account

import (
	documents_validate "BackEnd_Api/helpers/validations/document"
	email_validate "BackEnd_Api/helpers/validations/email"
	phone_validate "BackEnd_Api/helpers/validations/phone"
	roles_validate "BackEnd_Api/helpers/validations/roles"
	typeofaccounts_validate "BackEnd_Api/helpers/validations/typeofaccounts"
	account "BackEnd_Api/microservices/account/rules"
	"errors"
)

//
// microservices => account => usecase => update_account.go
//

func (aUC *AccountUseCase) UpdateAccount(a account.Account) (int, error) {

	err := email_validate.ValidateEmailFormat(a.Email)

	if err != nil {
		return 0, err
	}

	if !phone_validate.IsValidePhone(a.PhoneNumber, true, true, false) {
		return 0, errors.New("Telefone Formato Invalido")
	}

	if len(a.Document) > 0 {
		if !documents_validate.IsValidCPF(a.Document, true) && !documents_validate.IsValidCNPJ(a.Document, true) {
			return 0, errors.New("CPF ou CNPJ Invalido")
		}
	}

	if !roles_validate.IsValidRole(a.Roles...) {
		return 0, errors.New("Roles/Permissões Invalidos")
	}

	if !typeofaccounts_validate.IsValidTypeOfAccount(a.TypeOfAccount...) {
		return 0, errors.New("Tipo de Contas Invalidos")
	}

	i, err := aUC.repository.UpdateAccountRepository(a)

	if err != nil {
		return 0, err
	}

	return i, nil
}
