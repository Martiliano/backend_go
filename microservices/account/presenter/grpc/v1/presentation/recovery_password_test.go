package account

//
// microservices => account => presenter => grpc => v1 => presentation => recovery_password_test.go
//
//

import (
	"context"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"

	"BackEnd_Api/config"
	"BackEnd_Api/helpers/jwt"
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
)

func TestRecoveryPassword_Success(t *testing.T) {

	done := make(chan bool)
	quit := make(chan bool)

	var messageBody string

	go CreateNewAccountServer_Tester(done, quit, &messageBody)

	<-done

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, errDial := grpc.DialContext(ctx, "localhost:7070", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if errDial != nil {
		t.Error("não pode conectar ao servidor :", errDial)
	}
	defer conn.Close()

	cli := pb.NewAccountClient(conn)

	mockId, errMock := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "55988776655",
			Email:         "anyone1@email.com",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	})

	if errMock != nil {
		t.Error("Generate Account Password Token, Mock Error : ", errMock)
	}

	_, err := cli.GenerateRecoveryPasswordToken(ctx, &pb.GenerateRecoveryPasswordTokenRequest{Id: mockId.NewId})

	if err != nil {
		t.Error("Generate Account Password Token, Expected Success, Got :", err)
	}

	s1 := "<h1>RECUPERAÇÃO DE SENHA</h1><p></p><p>Abaixo segue o TOKEN necessário para você recuperar a sua senha no sistema.</p><p></p><h2>Token:</h2><p>"
	t1 := len(s1)
	s2 := messageBody[t1:]
	t2 := len(s2) - 4
	token := s2[:t2]

	_, err = cli.RecoveryPassword(ctx, &pb.RecoveryPasswordRequest{Token: token, Password: "Strong_PassWord",
		ConfirmPassword: "Strong_PassWord"})

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	if err != nil {
		t.Error("Generate Account Password Token, Expected ' Success ', Got :", err)
	}

}

func TestRecoveryPassword_InvalidToken(t *testing.T) {

	done := make(chan bool)
	quit := make(chan bool)

	var messageBody string

	go CreateNewAccountServer_Tester(done, quit, &messageBody)

	<-done

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, errDial := grpc.DialContext(ctx, "localhost:7070", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if errDial != nil {
		t.Error("não pode conectar ao servidor :", errDial)
	}
	defer conn.Close()

	cli := pb.NewAccountClient(conn)

	mockId, errMock := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "55988776655",
			Email:         "anyone1@email.com",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	})

	if errMock != nil {
		t.Error("Generate Account Password Token, Mock Error : ", errMock)
	}

	_, err := cli.GenerateRecoveryPasswordToken(ctx, &pb.GenerateRecoveryPasswordTokenRequest{Id: mockId.NewId})

	if err != nil {
		t.Error("Generate Account Password Token, Expected Success, Got :", err)
	}

	token := "yyyybGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY1NTM3MDgsImp0aSI6IjYyYmE1ZGVjNmViNWFlM2RkZTJmOWY0OSIsImlhdCI6MTY1NjM4MDkwOCwiaXNzIjoiaXVyaXMudjEucmVjb3ZlcnlwYXNzd29yZCIsInR5cGVfb2ZfYWNjb3VudCI6WyJEZXNlbnZvbHZlZG9yIl0sInJvbGVzIjpbIlNlbVJlc3RyaWNvZXMiLCJJbmNsdWlyIiwiRXhjbHVpciIsIkFsdGVyYXIiLCJDb25zdWx0YXIiLCJWaXN1YWxpemFyIl19.PDU7TdT0URRWeKtHyQIB5pCwdVmSkG-x3AvnO_JNxxx"

	_, err = cli.RecoveryPassword(ctx, &pb.RecoveryPasswordRequest{Token: token, Password: "Strong_PassWord",
		ConfirmPassword: "Strong_PassWord"})

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	msg := "rpc error: code = Internal desc = Token invalido"

	if err != nil && !strings.HasPrefix(err.Error(), msg) {
		t.Error("Generate Account Password Token, Expected ' rpc error: code = Internal desc = Token invalido ', Got :", err)
	}

}

func TestRecoveryPassword_NotExistToken(t *testing.T) {

	done := make(chan bool)
	quit := make(chan bool)

	var messageBody string

	go CreateNewAccountServer_Tester(done, quit, &messageBody)

	<-done

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, errDial := grpc.DialContext(ctx, "localhost:7070", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if errDial != nil {
		t.Error("não pode conectar ao servidor :", errDial)
	}
	defer conn.Close()

	cli := pb.NewAccountClient(conn)

	mock := pb.AccountEntities{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "55988776655",
		Email:         "anyone1@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	mockId, errMock := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &mock})

	if errMock != nil {
		t.Error("Generate Account Password Token, Mock Error : ", errMock)
	}

	_, err := cli.GenerateRecoveryPasswordToken(ctx, &pb.GenerateRecoveryPasswordTokenRequest{Id: mockId.NewId})

	if err != nil {
		t.Error("Generate Account Password Token, Expected Success, Got :", err)
	}

	config := config.GetConfig()
	jwtManager := jwt.NewJwtManager(config.Auth.CreateSecret, time.Hour*48)

	token, err := jwtManager.GenerateRecoveryPassword(mockId.NewId, mock.TypeOfAccount, []string{"Consultar", "Visualizar"})

	_, err = cli.RecoveryPassword(ctx, &pb.RecoveryPasswordRequest{Token: token, Password: "Strong_PassWord",
		ConfirmPassword: "Strong_PassWord"})

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	msg := "rpc error: code = Internal desc = mongo: no documents in result"

	if err != nil && !strings.HasPrefix(err.Error(), msg) {
		t.Error("Generate Account Password Token, Expected ' rpc error: code = Internal desc = Token invalido ', Got :", err)
	}

}
