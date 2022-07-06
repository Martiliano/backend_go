package account

//
// microservices => account => rules => entities.go
//

type Account struct {
	Id            string
	UserName      string
	PhoneNumber   string
	Email         string
	Document      string
	FullName      string
	TypeOfAccount []string
	Roles         []string
}
