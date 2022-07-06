package account

//
// microservices => account => presenter => grpc => v1 => presentation => create_account_test.go
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

func TestCreateAccount_Success(t *testing.T) {

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

	result, err := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
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

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done

	if err != nil {
		t.Error("Create Account By Secret, Expected : Success ", "Got :", err)
	}

	if !result.Success {
		t.Error("Create Account By Secret, Expected : Success ", "Got :", result.Success, " Message: ", result.Msg)
	}
}

func TestCreateAccount_InvalidPhoneNumber(t *testing.T) {

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

	result, err := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988s76655",
			Email:         "anyone@email.com",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	})

	msg := "Telefone Formato Invalido"

	if err == nil {
		t.Error("Create Account By Secret, Expected : '", msg, "' ' ", "Got :", result.Success, " Message: ", result.Msg)
	}

	if err != nil && !strings.HasSuffix(err.Error(), msg) {
		if result != nil {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got :", result.Success, " Message: ", result.Msg)
		} else {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got : ", err)
		}
	}

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done
}

func TestCreateAccount_InvalidEmail(t *testing.T) {

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

	result, err := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776655",
			Email:         "anyone@emailcom",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	})

	msg := "Email com Formato Invalido"

	if err == nil {
		t.Error("Create Account By Secret, Expected : '", msg, "' ' ", "Got :", result.Success, " Message: ", result.Msg)
	}

	if err != nil && !strings.HasSuffix(err.Error(), msg) {
		if result != nil {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got :", result.Success, " Message: ", result.Msg)
		} else {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got : ", err)
		}
	}

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done
}

func TestCreateAccount_InvalidDocument(t *testing.T) {

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

	result, err := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776655",
			Email:         "anyone@email.com",
			Document:      "00000000000",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	})

	msg := "CPF ou CNPJ Invalido"

	if err == nil {
		t.Error("Create Account By Secret, Expected : '", msg, "' ' ", "Got :", result.Success, " Message: ", result.Msg)
	}

	if err != nil && !strings.HasSuffix(err.Error(), msg) {
		if result != nil {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got :", result.Success, " Message: ", result.Msg)
		} else {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got : ", err)
		}
	}

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done
}

func TestCreateAccount_InvalidRoles(t *testing.T) {

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

	result, err := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776655",
			Email:         "anyone@email.com",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Blablabla", "Visualizar"},
		},
	})

	msg := "Roles/Permissões Invalidos"

	if err == nil {
		t.Error("Create Account By Secret, Expected : '", msg, "' ' ", "Got :", result.Success, " Message: ", result.Msg)
	}

	if err != nil && !strings.HasSuffix(err.Error(), msg) {
		if result != nil {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got :", result.Success, " Message: ", result.Msg)
		} else {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got : ", err)
		}
	}

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done
}

func TestCreateAccount_InvalidTypeOfAccounts(t *testing.T) {

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

	result, err := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776655",
			Email:         "anyone@email.com",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor", "Blabla"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	})

	msg := "Tipo de Contas Invalidos"

	if err == nil {
		t.Error("Create Account By Secret, Expected : '", msg, "' ' ", "Got :", result.Success, " Message: ", result.Msg)
	}

	if err != nil && !strings.HasSuffix(err.Error(), msg) {
		if result != nil {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got :", result.Success, " Message: ", result.Msg)
		} else {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got : ", err)
		}
	}

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done
}

func TestCreateAccount_ErrorEmailExists(t *testing.T) {

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

	mockId, errMock := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
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

	result, err := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
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

	msg := "rpc error: code = Internal desc = Erro já existe um usuário utilizando este e-mail"

	if err == nil {
		t.Error("Create Account By Secret, Expected : '", msg, "' ' ", "Got :", result.Success, " Message: ", result.Msg)
	}

	if err != nil && !strings.HasPrefix(err.Error(), msg) {
		if result != nil {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got :", result.Success, " Message: ", result.Msg)
		} else {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got : ", err)
		}
	}

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done
}

func TestCreateAccount_ErrorPhoneExists(t *testing.T) {

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

	mockId, errMock := cli.CreateAccount(ctx, &pb.CreateAccountRequest{
		Account: &pb.AccountEntities{
			Id:            "62956508bbde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776655",
			Email:         "anyone1@email.com",
			Document:      "04060664421",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	})

	result, err := cli.CreateAccountSecret(ctx, &pb.CreateAccountSecretRequest{
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

	msg := "rpc error: code = Internal desc = Erro já existe um usuário utilizando este telefone"

	if err == nil {
		t.Error("Create Account By Secret, Expected : '", msg, "' ' ", "Got :", result.Success, " Message: ", result.Msg)
	}

	if err != nil && !strings.HasPrefix(err.Error(), msg) {
		if result != nil {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got :", result.Success, " Message: ", result.Msg)
		} else {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got : ", err)
		}
	}

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done
}

func TestCreateAccount_ErrorDocumentExists(t *testing.T) {

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

	result, err := cli.CreateAccountSecret(ctx, &pb.CreateAccountSecretRequest{
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

	msg := "rpc error: code = Internal desc = Erro já existe um usuário utilizando este documento"

	if err == nil {
		t.Error("Create Account By Secret, Expected : '", msg, "' ' ", "Got :", result.Success, " Message: ", result.Msg)
	}

	if err != nil && !strings.HasPrefix(err.Error(), msg) {
		if result != nil {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got :", result.Success, " Message: ", result.Msg)
		} else {
			t.Error("Create Account By Secret, Expected : '", msg, "' ", "Got : ", err)
		}
	}

	if errMock == nil && mockId.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: mockId.NewId})
	}

	if err == nil && result.Success {
		_, _ = cli.DeleteAccount(ctx, &pb.DeleteAccountRequest{Id: result.NewId})
	}

	quit <- true

	<-done
}
