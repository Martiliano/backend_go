package auth

//
// microservices => account => usecase => login_by_email_test.go
//

import (
	config "BackEnd_Api/config"
	account_repository_mongo "BackEnd_Api/microservices/account/external/db/mongo/implementation"
	"strings"

	"BackEnd_Api/logger"
	"context"

	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func TestLoginByEmail_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		t.Error("Login By Email, Expected JWT Token, Got :", err)
		return
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		t.Error("Login By Email, Expected JWT Token, Got :", err)
		return
	}

	repo_account := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")
	config := config.GetConfig()

	passwordAccountStored := "Strong_Password"
	passwordForValidate := "Strong_Password"

	repo_auth := NewMockAuthRepository(passwordAccountStored)

	isSucesso = true

	usecase := NewAuthUseCase(repo_auth, repo_account, config)

	_, err = usecase.LoginByEmail("email@email.com", passwordForValidate)

	if err != nil {
		t.Error("Login By Email, Expected JWT Token, Got :", err)
		return
	}

}

func TestLoginByEmail_IncorrectPassword(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		t.Error("Login By Email, Expected JWT Token, Got :", err)
		return
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		t.Error("Login By Email, Expected JWT Token, Got :", err)
		return
	}

	repo_account := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")
	config := config.GetConfig()

	passwordAccountStored := "Strong_Password"
	passwordForValidate := "Incorrect_Password"

	repo_auth := NewMockAuthRepository(passwordAccountStored)

	usecase := NewAuthUseCase(repo_auth, repo_account, config)

	isSucesso = true

	token, err := usecase.LoginByEmail("email@email.com", passwordForValidate)

	if err == nil {
		t.Error("Login By Email, Expected ERROR, Got :", token)
		return
	}

	if err != nil && !strings.HasPrefix(err.Error(), "crypto/bcrypt: hashedPassword is not the hash of the given password") {
		t.Error("Login By Email, Expected ERROR 'crypto/bcrypt: hashedPassword is not the hash of the given password', Got :", err)
		return
	}
	//
}

func TestLoginByEmail_EmailNotExists(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		t.Error("Login By Email, Expected JWT Token, Got :", err)
		return
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		t.Error("Login By Email, Expected JWT Token, Got :", err)
		return
	}

	repo_account := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")
	config := config.GetConfig()

	passwordAccountStored := "Strong_Password"
	passwordForValidate := "Incorrect_Password"

	repo_auth := NewMockAuthRepository(passwordAccountStored)

	usecase := NewAuthUseCase(repo_auth, repo_account, config)

	isSucesso = false

	token, err := usecase.LoginByEmail("email@email.com", passwordForValidate)

	if err == nil {
		t.Error("Login By Email, Expected ERROR, Got :", token)
		return
	}

	if err != nil && !strings.HasPrefix(err.Error(), "Não existe a Account referida pelo Email") {
		t.Error("Login By Email, Expected ERROR 'Não existe a Account referida pelo Email', Got :", err)
		return
	}

}
