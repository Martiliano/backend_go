package account

//
// microservices => account => usecase => get_account_by_email.go
//

import (
	account "BackEnd_Api/microservices/account/rules"
)

func (aUC *AccountUseCase) GetAccountByEmail(email string) (*account.Account, error) {

	a, err := aUC.repository.GetAccountByEmailRepository(email)

	if err != nil {
		return nil, err
	}

	return a, nil

}
