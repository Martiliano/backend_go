package auth

//
// microservices => auth => rules => contracts.go
//

type IAuth interface {
	LoginById(id string, password string) (string, error)
	LoginByEmail(email string, password string) (string, error)
	LoginByPhone(phone string, password string) (string, error)
	LoginByDocument(document string, password string) (string, error)
	VersionOfAuthMicroService() string
}
