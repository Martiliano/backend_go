package auth_repository

//
// microservices => auth => repository => contracts.go
//

type IAuthRepository interface {
	GetAccountPasswordHashByIdRepository(id string) (string, error)
	GetAccountPasswordHashByEmailRepository(email string) (string, error)
	GetAccountPasswordHashByDocumentRepository(document string) (string, error)
	GetAccountPasswordHashByPhoneRepository(phone string) (string, error)
}
