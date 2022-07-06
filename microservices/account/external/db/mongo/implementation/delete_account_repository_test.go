package account

//
// microservices => account => external => db => mongo => implementation => delete_account_repository_test.go
//

import (
	account "BackEnd_Api/microservices/account/rules"

	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeleteAccountRepository(t *testing.T) {
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
		Email:         "anyone@email.com",
		Document:      "12345678900",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	repository := NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")

	id, err := repository.CreateAccountRepository(a)

	if err != nil {
		t.Error("Delete Account, Error Create Test Account :", err)
	}

	count, err := repository.DeleteAccountRepository(id)

	if err != nil {
		t.Error("Delete Account, Expected deleted count 1, Got :", err)
	}

	if count != 1 {
		t.Error("Delete Account, Expected deleted count 1, Got :", count)
	}

}
