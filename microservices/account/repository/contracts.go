package account_repository

//
// microservices => account => repository => contracts.go
//
// Repository - Interfaces
//

import (
	account "BackEnd_Api/microservices/account/rules"
)

type IAccountRepository interface {
	CreateAccountRepository(a account.Account) (string, error)

	GetAccountByIdRepository(id string) (*account.Account, error)
	GetAccountByEmailRepository(email string) (*account.Account, error)
	GetAccountByPhoneRepository(phone string) (*account.Account, error)
	GetAccountByDocumentRepository(document string) (*account.Account, error)

	GetAllAccountsRepository() (*[]account.Account, error)
	UpdateAccountRepository(a account.Account) (int, error)
	DeleteAccountRepository(id string) (int, error)

	GetAccountBySetPasswordTokenRepository(token string) (*account.Account, error)
	GetAccountByRecoveryTokenRepository(token string) (*account.Account, error)

	GenerateSetPasswordTokenRepository(id string, token string) error
	SetPasswordRepository(id string, passwordHash string) (int, error)

	GenerateRecoveryPasswordTokenRepository(id string, token string) error
	RecoveryPasswordRepository(id string, passwordHash string) (int, error)
}
