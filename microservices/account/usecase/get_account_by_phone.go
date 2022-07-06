package account

//
// microservices => account => usecase => get_account_by_phone.go
//

import (
	account "BackEnd_Api/microservices/account/rules"
)

func (aUC *AccountUseCase) GetAccountByPhone(phone string) (*account.Account, error) {

	a, err := aUC.repository.GetAccountByPhoneRepository(phone)

	if err != nil {
		return nil, err
	}

	return a, nil

}
