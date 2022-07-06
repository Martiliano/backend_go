package auth

//
// microservices => account => usecase => version_of_auth_microservice.go
//

func (aUC *AuthUseCase) VersionOfAuthMicroService() string {

	return "0.1.0-Auth"
}
