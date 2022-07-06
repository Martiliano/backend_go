package account

//
// microservices => account => usecase => get_all_accounts.go
//

import (
	account "BackEnd_Api/microservices/account/rules"
)

func (aUC *AccountUseCase) GetAllAccounts() (*[]account.Account, error) {

	a, err := aUC.repository.GetAllAccountsRepository()

	if err != nil {
		return nil, err
	}

	return a, nil
}
