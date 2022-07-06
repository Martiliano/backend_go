package account

//
// microservices => account => presenter => grpc => v1 => presentation => get_account_by_document_test.go
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

func TestGetAccountByDocument_Success(t *testing.T) {

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
		t.Error("Generate Account Password Token, Mock Error : ", errMock)
	}

	a, err := cli.GetAccountByDocument(ctx, &pb.GetAccountByDocumentRequest{Document: mock.Account.Document})

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	if err != nil {
		t.Error("Get Account By Document, Expected Account by document ' ", mock.Account.Document, "', Got :", err)
	} else if mock.Account.Document != a.Account.Document {
		t.Error("Get Account By Document, Expected Account by document ' ", mock.Account.Document, "', Got :", a.Account.Document)
	}

}

func TestGetAccountByDocument_NotExistDocument(t *testing.T) {

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

	fakeDoc := "11111111111"

	mockId, errMock := cli.CreateAccount(ctx, &mock)

	if errMock != nil {
		t.Error("Generate Account Password Token, Mock Error : ", errMock)
	}

	a, err := cli.GetAccountByDocument(ctx, &pb.GetAccountByDocumentRequest{Document: fakeDoc})

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	msg := "rpc error: code = Internal desc = mongo: no documents in result"

	if err != nil {
		if !strings.HasPrefix(err.Error(), msg) {
			t.Error("Get Account By Document, Expected Error ' ", msg, "', Got :", err)
		}
	} else if mock.Account.Document != a.Account.Document {
		t.Error("Get Account By Document, Expected Account by document ' ", mock.Account.Document, "', Got :", a.Account.Document)
	}

}
