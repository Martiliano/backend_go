package account

//
// microservices => account => external => db => mongo => implementation => get_account_by_phone_repository_test.go
//

import (
	account "BackEnd_Api/microservices/account/rules"

	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetAccountByPhoneRepository(t *testing.T) {
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
		PhoneNumber:   "54987651234",
		Email:         "anyone4544@email.com",
		Document:      "12315678949",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	repository := NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")

	id, err := repository.CreateAccountRepository(a)

	if err != nil {
		t.Error("GetAccountByPhone Account, Error Create Test Account :", err)
	}

	result, err := repository.GetAccountByPhoneRepository(a.PhoneNumber)

	repository.DeleteAccountRepository(id)

	if err != nil {
		t.Error("GetAccountByPhone Account, Error Create Test Account :", err)
	}

	if id != result.Id {
		t.Error("GetAccountByPhone Account, Expected :", id, ", Got : ", result.Id)
	}
}
