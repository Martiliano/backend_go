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

func TestUpdateAccount_Success(t *testing.T) {

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

	mockId, errMock := cli.CreateAccount(ctx, &pb.CreateAccountRequest{Account: &mock})

	if errMock != nil {
		t.Error("Mock Error : ", errMock)
	}

	mock.Id = mockId.NewId
	mock.FullName = "Anyone Update"

	count, err := cli.UpdateAccount(ctx, &pb.UpdateAccountRequest{Account: &mock})

	if err != nil {
		t.Error("Update Account, Expected ' Success ', Got :", err)
	}

	updated, err := cli.GetAccountById(ctx, &pb.GetAccountByIdRequest{Id: mockId.NewId})

	if err != nil {
		t.Error("Update Account, Expected ' Success ', Got :", err)
	}

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	if updated.Account.FullName != mock.FullName {
		t.Error("Update Account, Expected ' ", mock.FullName, " ', Got :", updated.Account.FullName)
	}

	if count.NumberOfUpdates != 1 {
		t.Error("Update Account, Expected ' 1 ', Got :", count.NumberOfUpdates)
	}

}

func TestUpdateAccount_NotExistId(t *testing.T) {

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

	mockId, errMock := cli.CreateAccount(ctx, &pb.CreateAccountRequest{Account: &mock})

	fakeId := "52956508bbde2d63320a8a49"

	if errMock != nil {
		t.Error("Mock Error : ", errMock)
	}

	mock.Id = fakeId
	mock.FullName = "Anyone Update"

	count, err := cli.UpdateAccount(ctx, &pb.UpdateAccountRequest{Account: &mock})

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	quit <- true

	<-done

	msg := "rpc error: code = Internal desc = Registro a ser alterado não encontrado"

	if err != nil {
		if !strings.HasPrefix(err.Error(), msg) {
			t.Error("Get Account, Expected Error ' ", msg, "', Got :", err)
		}
	} else if count.NumberOfUpdates != 0 {
		t.Error("Update Account, Expected ' 0 ', Got :", count.NumberOfUpdates)
	}

}
