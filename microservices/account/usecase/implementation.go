package account

//
// microservices => account => usecase => implementation.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email"
	account_repository "BackEnd_Api/microservices/account/repository"
	account "BackEnd_Api/microservices/account/rules"
)

type AccountUseCase struct {
	config       *config.Config
	repository   account_repository.IAccountRepository
	email        email_service.IEmail
	serverSecret string
}

func NewAccountUseCase(r account_repository.IAccountRepository, e email_service.IEmail, c *config.Config, s string) account.IAccount {
	return &AccountUseCase{repository: r, serverSecret: s, config: c, email: e}
}
