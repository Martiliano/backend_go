package account

//
// microservices => account => usecase => delete_account_test.go
//

import (
	config "BackEnd_Api/config"
	"BackEnd_Api/logger"
	account "BackEnd_Api/microservices/account/rules"
	"strings"

	"testing"

	"go.uber.org/zap"
)

func TestDeleteAccount_Success(t *testing.T) {
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	message := make(chan string)
	defer close(message)

	done := make(chan bool)
	defer close(done)

	repo := NewMockAccountRepository("Token")
	email := NewEmailMockService(message)
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776655",
		Email:         "anyone@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	count, err := usecase.UpdateAccount(a)

	if err != nil {
		t.Error("Update Account, Expected update Account Id :", a.Id, "Got :", err)
	} else if count == 0 {
		t.Error("Update Account, Expected update Account Id :", a.Id, "Got :", count)
	}
}

func TestDeleteAccount_Error(t *testing.T) {
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	message := make(chan string)
	defer close(message)

	done := make(chan bool)
	defer close(done)

	repo := NewMockAccountRepository("Token")
	email := NewEmailMockService(message)
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776655",
		Email:         "anyone@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = false

	_, err := usecase.UpdateAccount(a)
	msg := "Erro ao Atualizar Account"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}
