package account

//
// microservices => account => usecase => get_all_accounts_test_test.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	"BackEnd_Api/logger"
	"testing"

	"go.uber.org/zap"
)

func TestGetAllAccounts_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAllAccounts()

	if err != nil {
		t.Error("Get All Account, Expected Account list 3 elements, ", "Got :", err)
	} else if len(*a) != 3 {
		t.Error("Get All Account, Expected Account list 3 elements, ", "Got :", len(*a))
	}
}

func TestGetAllAccounts_Error(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = false

	a, err := usecase.GetAllAccounts()
	msg := "Erro ao listar todas as Account."

	if err == nil {
		t.Error("Get Account by PhoneNumber, Expected ERROR", ", Got list of ", len(*a), " elements")
	} else if err.Error() != msg {
		t.Error("Get Account by PhoneNumber, Expected error message '", msg, "', Got:", err)
	}

}
