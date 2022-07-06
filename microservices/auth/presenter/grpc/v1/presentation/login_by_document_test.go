package auth

//
// microservices => auth => presenter => grpc => v1 => presentation => login_by_document_test.go
//
//

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"

	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	account_repository_mongo "BackEnd_Api/microservices/account/external/db/mongo/implementation"
	account "BackEnd_Api/microservices/account/rules"
	account_usecase "BackEnd_Api/microservices/account/usecase"
	pb "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"
)

func TestLoginByDocument_Success(t *testing.T) {
	done := make(chan bool)
	quit := make(chan bool)

	go CreateNewAuthServer_Tester(done, quit)

	<-done

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, errDial := grpc.DialContext(ctx, "localhost:7070", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if errDial != nil {
		t.Error("não pode conectar ao servidor :", errDial)
	}
	defer conn.Close()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		log.Printf("Erro ao conectar ao MongoDb : " + error.Error(err))
		return
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		log.Printf("Erro ao testar a conexão ao MongoDb : " + error.Error(err))
		return
	}

	config := config.GetConfig()
	email := email_service.NewEmailService()
	email.Initialize()

	smtpSecure := false
	if config.Email.SmtpSecure == "ssl/tls" {
		smtpSecure = true
	}

	smtpPort, _ := strconv.Atoi(config.Email.SmtpPort)

	email.SetSmtpServer(config.Email.SmtpServer, smtpSecure, smtpPort)
	email.SetEmailCredentials(config.Email.User, config.Email.Password)
	email.SendEmailFrom(config.Email.Account, config.Email.AccountName)

	account_repository_mongo := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")

	account_usecase := account_usecase.NewAccountUseCase(account_repository_mongo, email, config, config.Auth.CreateSecret)

	cli := pb.NewAuthClient(conn)

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "55988776655",
		Email:         "anyone1@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	id, errMock := account_usecase.CreateAccount(a)

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got :", err)
	}

	password := "StrongPassword"

	hashPassord, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got :", err)
	}

	_, err = account_repository_mongo.SetPasswordRepository(id, string(hashPassord))

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got :", err)
	}

	_, err = cli.LoginByDocument(ctx, &pb.LoginByDocumentRequest{Document: a.Document, Password: password})

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got :", err)
	}

	if errMock == nil {
		_, _ = account_usecase.DeleteAccount(id)
	}

	quit <- true

	<-done

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got :", err)
	}

}

func TestLoginByDocument_DocumentNotExists(t *testing.T) {

	done := make(chan bool)
	quit := make(chan bool)

	go CreateNewAuthServer_Tester(done, quit)

	<-done

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, errDial := grpc.DialContext(ctx, "localhost:7070", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if errDial != nil {
		t.Error("não pode conectar ao servidor (1) :", errDial)
	}
	defer conn.Close()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		log.Printf("Erro ao conectar ao MongoDb (2) : " + error.Error(err))
		return
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		log.Printf("Erro ao testar a conexão ao MongoDb (3) : " + error.Error(err))
		return
	}

	config := config.GetConfig()
	email := email_service.NewEmailService()
	email.Initialize()

	smtpSecure := false
	if config.Email.SmtpSecure == "ssl/tls" {
		smtpSecure = true
	}

	smtpPort, _ := strconv.Atoi(config.Email.SmtpPort)

	email.SetSmtpServer(config.Email.SmtpServer, smtpSecure, smtpPort)
	email.SetEmailCredentials(config.Email.User, config.Email.Password)
	email.SendEmailFrom(config.Email.Account, config.Email.AccountName)

	account_repository_mongo := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")

	account_usecase := account_usecase.NewAccountUseCase(account_repository_mongo, email, config, config.Auth.CreateSecret)

	cli := pb.NewAuthClient(conn)

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "55988776655",
		Email:         "anyone1@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	id, errMock := account_usecase.CreateAccount(a)

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got (4) :", err)
	}

	fmt.Println("TestLoginByDocument_DocumentNotExists id (5) : ", id)

	password := "StrongPassword"

	hashPassord, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got (6) :", err)
	}

	_, err = account_repository_mongo.SetPasswordRepository(id, string(hashPassord))

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got (7) :", err)
	}

	msg := "rpc error: code = Unknown desc = mongo: no documents in result"

	_, err = cli.LoginByDocument(ctx, &pb.LoginByDocumentRequest{Document: "66404617811", Password: password})

	if err != nil && !strings.HasPrefix(err.Error(), msg) {
		t.Error("Login By Document, Expected  PasswordHash, Got (8) :", err)
	}

	if errMock == nil {
		_, _ = account_usecase.DeleteAccount(id)
	}

	quit <- true

	<-done
}

func TestLoginByDocument_IncorrectPassword(t *testing.T) {

	done := make(chan bool)
	quit := make(chan bool)

	go CreateNewAuthServer_Tester(done, quit)

	<-done

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, errDial := grpc.DialContext(ctx, "localhost:7070", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if errDial != nil {
		t.Error("não pode conectar ao servidor (1) :", errDial)
	}
	defer conn.Close()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		log.Printf("Erro ao conectar ao MongoDb (2) : " + error.Error(err))
		return
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		log.Printf("Erro ao testar a conexão ao MongoDb (3) : " + error.Error(err))
		return
	}

	config := config.GetConfig()
	email := email_service.NewEmailService()
	email.Initialize()

	smtpSecure := false
	if config.Email.SmtpSecure == "ssl/tls" {
		smtpSecure = true
	}

	smtpPort, _ := strconv.Atoi(config.Email.SmtpPort)

	email.SetSmtpServer(config.Email.SmtpServer, smtpSecure, smtpPort)
	email.SetEmailCredentials(config.Email.User, config.Email.Password)
	email.SendEmailFrom(config.Email.Account, config.Email.AccountName)

	account_repository_mongo := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")

	account_usecase := account_usecase.NewAccountUseCase(account_repository_mongo, email, config, config.Auth.CreateSecret)

	cli := pb.NewAuthClient(conn)

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "55988776655",
		Email:         "anyone1@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	id, errMock := account_usecase.CreateAccount(a)

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got (4) :", err)
	}

	password := "StrongPassword"

	hashPassord, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got (6) :", err)
	}

	_, err = account_repository_mongo.SetPasswordRepository(id, string(hashPassord))

	if err != nil {
		t.Error("Login By Document, Expected  PasswordHash, Got (7) :", err)
	}

	msg := "rpc error: code = Unknown desc = crypto/bcrypt: hashedPassword is not the hash of the given password"

	_, err = cli.LoginByDocument(ctx, &pb.LoginByDocumentRequest{Document: a.Document, Password: "IncorrectPassword"})

	if err != nil && !strings.HasPrefix(err.Error(), msg) {
		t.Error("Login By Document, Expected  PasswordHash, Got (8) :", err)
	}

	if errMock == nil {
		_, _ = account_usecase.DeleteAccount(id)
	}

	quit <- true

	<-done
}
