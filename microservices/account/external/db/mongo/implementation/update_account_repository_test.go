//update_account_repository_test.go
package account

//
// microservices => account => external => db => mongo => implementation => set_password_repository_test.go
//

import (
	account "BackEnd_Api/microservices/account/rules"

	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestUpdateAccountRepository(t *testing.T) {
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
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	repository := NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")

	id, err := repository.CreateAccountRepository(a)

	if err != nil {
		t.Error("UpdateAccount Account, Error Create Test Account :", err)
	}

	a.Id = id
	a.FullName = "Anyone Unknown Alter"

	result, err := repository.UpdateAccountRepository(a)

	repository.DeleteAccountRepository(id)

	if err != nil {
		t.Error("UpdateAccount Account, Error Create Test Account :", err)
	}

	if result != 1 {
		t.Error("UpdateAccount Account, Expected : 1, Got : ", result)
	}
}
