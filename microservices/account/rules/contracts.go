package account

//
// microservices => account => rules => contracts.go
//

type IAccount interface {
	CreateAccountBySecret(a Account, secret string) (string, error)
	VersionOfAccountMicroService() string

	CreateAccount(a Account) (string, error)

	GenerateSetPasswordToken(id string) error
	SetPassword(token string, password string, confirmPassword string) error

	GenerateRecoveryPasswordToken(id string) error
	RecoveryPassword(token string, password string, confirmPassword string) error

	GetAccountById(id string) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
	GetAccountByPhone(phone string) (*Account, error)
	GetAccountByDocument(document string) (*Account, error)

	GetAllAccounts() (*[]Account, error)

	UpdateAccount(a Account) (int, error)

	DeleteAccount(id string) (int, error)
}
