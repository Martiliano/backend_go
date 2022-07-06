package account

//
// microservices => account => usecase => delete_account.go
//

func (aUC *AccountUseCase) DeleteAccount(id string) (int, error) {

	i, err := aUC.repository.DeleteAccountRepository(id)

	if err != nil {
		return 0, err
	}

	return i, nil
}
