package account

//
// microservices => account => usecase => version_of_account_microservice.go
//

func (aUC *AccountUseCase) VersionOfAccountMicroService() string {
	// Versão.
	return "0.1.0-Account"
}
