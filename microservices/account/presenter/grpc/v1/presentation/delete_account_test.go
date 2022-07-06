package account

//
// microservices => account => presenter => grpc => v1 => presentation => delete_account_test.go
//
//

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
)

func TestDeleteAccount_Success(t *testing.T) {

	done := make(chan bool)
	quit := make(chan bool)

	var messageBody string

	go CreateNewAccountServer_Tester(done, quit, &messageBody)

	<-done

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost:7070", grpc.WithBlock(), grpc.WithTimeout(5*time.Second), grpc.WithInsecure())
	if err != nil {
		t.Error("não pode conectar ao servidor :", err)
	}
	defer conn.Close()

	cli := pb.NewAccountClient(conn)

	createResult, err := cli.CreateAccountSecret(ctx, &pb.CreateAccountSecretRequest{
		Secret: "secret_create_phrase",
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776655",
			Email:         "anyone@email.com",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	})

	if err != nil {
		t.Error("Create Account By Secret, Expected : Success ", "Got :", err)
	}

	result, err := cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: createResult.NewId})

	quit <- true

	<-done

	if err != nil {
		t.Error("Create Account By Secret, Expected : Success ", "Got :", err)
	}

	if !result.Success {
		t.Error("Create Account By Secret, Expected : Success ", "Got :", result.Success, " Message: ", result.Msg)
	}
}
