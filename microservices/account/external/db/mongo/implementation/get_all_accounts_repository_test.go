package account

//
// microservices => account => external => db => mongo => implementation => get_all_accounts_repository_test.go
//

import (
	account "BackEnd_Api/microservices/account/rules"

	"fmt"

	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetAllAccountsRepository(t *testing.T) {
	ctx := context.Background()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		t.Error("Create Account, Expected new Account Id : - , Got :", err)
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		t.Error("Create Account, Expected new Account Id : - , Got :", err)
	}

	a := account.Account{
		UserName:      "anyone",
		PhoneNumber:   "54988776655",
		Email:         "anyone4544@email.com",
		Document:      "12315678949",
		FullName:      "Anyone 1 Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	repository := NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")

	id1, err := repository.CreateAccountRepository(a)

	a.FullName = "Anyone 2 Unknown"
	id2, err := repository.CreateAccountRepository(a)

	a.FullName = "Anyone 3 Unknown"
	id3, err := repository.CreateAccountRepository(a)

	if err != nil {
		t.Error("GetAccountById Account, Error Create Test Account :", err)
	}

	results, err := repository.GetAllAccountsRepository()

	fmt.Println("results[0].Id : ", (*results)[0].Id)
	fmt.Println("len((*results)) : ", len((*results)))

	repository.DeleteAccountRepository(id1)
	repository.DeleteAccountRepository(id2)
	repository.DeleteAccountRepository(id3)

	if err != nil {
		t.Error("GetAccountById Account, Error Create Test Account :", err)
	}

	if len((*results)) != 3 {
		t.Error("GetAccountById Account, Expected :", 3, ", Got : ", len((*results)))
	}
}
