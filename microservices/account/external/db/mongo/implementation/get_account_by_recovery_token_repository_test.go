package account

//
// microservices => account => external => db => mongo => implementation => get_account_by_recovery_token_repository_test.go
//

import (
	"BackEnd_Api/helpers/jwt"
	account "BackEnd_Api/microservices/account/rules"
	"time"

	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetAccountByRecoveryTokenRepository(t *testing.T) {
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
		t.Error("GetAccountByRecoveryToken Account, Error Create Test Account :", err)
	}

	jwtManager := jwt.NewJwtManager("Secret_Phrase", time.Hour*48)

	token, err := jwtManager.GenerateRecoveryPassword(id, a.TypeOfAccount, a.Roles)

	if err != nil {
		t.Error("GetAccountByRecoveryToken Account, Error Generate Token :", err)
	}

	err = repository.GenerateRecoveryPasswordTokenRepository(id, token)

	if err != nil {
		t.Error("GetAccountByRecoveryToken Account, Error Save Generate Token :", err)
	}

	result, err := repository.GetAccountByRecoveryTokenRepository(token)

	repository.DeleteAccountRepository(id)

	if err != nil {
		t.Error("GetAccountByRecoveryToken Account, Error Get Account By Code :", err)
	}

	if id != result.Id {
		t.Error("GetAccountByRecoveryToken Account, Expected :", id, ", Got : ", result.Id)
	}
}
