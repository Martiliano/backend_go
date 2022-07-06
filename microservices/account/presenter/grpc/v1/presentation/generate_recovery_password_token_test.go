package account

//
// microservices => account => presenter => grpc => v1 => presentation => generate_recovery_password_token_test.go
//
//

import (
	"context"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
)

func TestGenerateRecoveryPasswordToken_Success(t *testing.T) {

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

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	result := messageBody
	msg := "<h1>RECUPERAÇÃO DE SENHA</h1>"

	quit <- true

	<-done

	if err != nil {
		t.Error("Generate Account Password Token, Expected ' ", msg, "', Got :", err)
	} else if !strings.HasPrefix(result, msg) {
		t.Error("Generate Account Password Token, Expected ' ", msg, "', Got :", result)
	}

}
