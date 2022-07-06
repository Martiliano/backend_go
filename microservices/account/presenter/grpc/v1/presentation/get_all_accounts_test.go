package account

//
// microservices => account => presenter => grpc => v1 => presentation => get_all_accounts_test.go
//
//

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
)

func TestGetAllAccounts_Success(t *testing.T) {

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

	mock1 := pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone1",
			PhoneNumber:   "55988776655",
			Email:         "anyone1@email.com",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	}

	mock2 := pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone1",
			PhoneNumber:   "55988776656",
			Email:         "anyone2@email.com",
			Document:      "45721236400",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	}

	mock3 := pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone1",
			PhoneNumber:   "55988776657",
			Email:         "anyone3@email.com",
			Document:      "66671871418",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	}

	mockId1, errMock1 := cli.CreateAccount(ctx, &mock1)

	if errMock1 != nil {
		t.Error("Mock Error : ", errMock1)
	}

	mockId2, errMock2 := cli.CreateAccount(ctx, &mock2)

	if errMock2 != nil {
		t.Error("Mock Error : ", errMock2)
	}

	mockId3, errMock3 := cli.CreateAccount(ctx, &mock3)

	if errMock3 != nil {
		t.Error("Mock Error : ", errMock3)
	}

	a, err := cli.GetAllAccounts(ctx, &pb.GetAllAccountsRequest{})

	if errMock1 == nil && mockId1.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId1.NewId})
	}

	if errMock2 == nil && mockId2.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId2.NewId})
	}

	if errMock3 == nil && mockId2.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId3.NewId})
	}

	quit <- true

	<-done

	if err != nil {
		t.Error("Get All Account, Expected 3 Itens, Got :", err)
	} else if len(a.Account) != 3 {
		t.Error("Get All Account, Expected 3 Itens, Got :", len(a.Account))
	}

}
