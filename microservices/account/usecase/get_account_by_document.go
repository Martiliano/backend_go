package account

//
// microservices => account => usecase => get_account_by_document.go
//

import (
	account "BackEnd_Api/microservices/account/rules"
)

func (aUC *AccountUseCase) GetAccountByDocument(document string) (*account.Account, error) {

	a, err := aUC.repository.GetAccountByDocumentRepository(document)

	if err != nil {
		return nil, err
	}

	return a, nil

}
