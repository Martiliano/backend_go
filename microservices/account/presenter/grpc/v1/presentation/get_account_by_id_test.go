package account

//
// microservices => account => presenter => grpc => v1 => presentation => get_account_by_id_test.go
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

func TestGetAccountById_Success(t *testing.T) {

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

	mock := pb.CreateAccountRequest{
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
	}

	mockId, errMock := cli.CreateAccount(ctx, &mock)

	if errMock != nil {
		t.Error("Mock Error : ", errMock)
	}

	a, err := cli.GetAccountById(ctx, &pb.GetAccountByIdRequest{Id: mockId.NewId})

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	if err != nil {
		t.Error("Get Account By Id, Expected Account by Id ' ", mock.Account.Id, "', Got :", err)
	} else if mockId.NewId != a.Account.Id {
		t.Error("Get Account By Id, Expected Account by Id ' ", mock.Account.Id, "', Got :", a.Account.Id)
	}

}

func TestGetAccountById_NotExistId(t *testing.T) {

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

	mock := pb.CreateAccountRequest{
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
	}

	fakeId := "52956508bbde2d63320a8a49"

	mockId, errMock := cli.CreateAccount(ctx, &mock)

	if errMock != nil {
		t.Error("Mock Error : ", errMock)
	}

	a, err := cli.GetAccountById(ctx, &pb.GetAccountByIdRequest{Id: fakeId})

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	msg := "rpc error: code = Internal desc = mongo: no documents in result"

	if err != nil {
		if !strings.HasPrefix(err.Error(), msg) {
			t.Error("Get Account By Id, Expected Error ' ", msg, "', Got :", err)
		}
	} else if mockId.NewId != a.Account.Id {
		t.Error("Get Account By Id, Expected Account by Id ' ", mock.Account.Id, "', Got :", a.Account.Id)
	}

}
