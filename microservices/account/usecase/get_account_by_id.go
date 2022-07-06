package account

//
// microservices => account => usecase => get_account_by_id.go
//

import (
	account "BackEnd_Api/microservices/account/rules"
)

func (aUC *AccountUseCase) GetAccountById(id string) (*account.Account, error) {

	a, err := aUC.repository.GetAccountByIdRepository(id)

	if err != nil {
		return nil, err
	}

	return a, nil
}
