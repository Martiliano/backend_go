package auth

//
// microservices => auth => external => db => mongo => implementation => get_account_password_hash_by_document_test.go
//

import (
	account_repository_mongo "BackEnd_Api/microservices/account/external/db/mongo/implementation"
	account "BackEnd_Api/microservices/account/rules"
	"strings"

	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func TestGetAccountPasswordHashByDocumentRepository_Success(t *testing.T) {
	ctx := context.Background()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
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

	repository_account := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")
	repository_auth := NewAuthMongoRepository(db, ctx, "BackEnd_Api_test")

	id, err := repository_account.CreateAccountRepository(a)

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
	}

	hashPassord, err := bcrypt.GenerateFromPassword([]byte("SecretPhrase"), 14)

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
	}

	_, err = repository_account.SetPasswordRepository(id, string(hashPassord))

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
	}

	_, err = repository_auth.GetAccountPasswordHashByDocumentRepository(a.Document)

	repository_account.DeleteAccountRepository(id)

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
	}
}

func TestGetAccountPasswordHashByDocumentRepository_InvalidHashPassword(t *testing.T) {
	ctx := context.Background()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
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

	repository_account := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")
	repository_auth := NewAuthMongoRepository(db, ctx, "BackEnd_Api_test")

	id, err := repository_account.CreateAccountRepository(a)

	if err != nil {
		t.Error("Get Account PasswordHash By Document, Expected  PasswordHash, Got :", err)
	}

	_, err = repository_auth.GetAccountPasswordHashByDocumentRepository(a.Document)

	repository_account.DeleteAccountRepository(id)

	msg := "Password Hash Invalida"

	if err != nil && !strings.HasPrefix(err.Error(), msg) {
		t.Error("Get Account PasswordHash By Document, Expected ' ", msg, " ', Got :", err)
	}
}
