package auth

//
// microservices => auth => usecase => contracts.go
//

import (
	config "BackEnd_Api/config"
	account_repository "BackEnd_Api/microservices/account/repository"
	auth_repository "BackEnd_Api/microservices/auth/repository"
	auth "BackEnd_Api/microservices/auth/rules"
)

type AuthUseCase struct {
	config             *config.Config
	account_repository account_repository.IAccountRepository
	auth_repository    auth_repository.IAuthRepository
}

func NewAuthUseCase(r auth_repository.IAuthRepository, a account_repository.IAccountRepository, c *config.Config) auth.IAuth {
	return &AuthUseCase{auth_repository: r, account_repository: a, config: c}
}
