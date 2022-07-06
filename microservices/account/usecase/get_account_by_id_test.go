package account

//
// microservices => account => usecase => get_account_by_id_test.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	"BackEnd_Api/logger"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func TestGetAccountById_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAccountById("62956508ccde2d63320a8a59")

	if err != nil {
		t.Error("Get Account by Id, Expected Account Id :", "62956508ccde2d63320a8a59", "Got :", err)
	} else if a.Id != "62956508ccde2d63320a8a59" {
		t.Error("Get Account by Id, Expected Account Id :", "62956508ccde2d63320a8a59", "Got :", a.Id)
	}
}

func TestGetAccountById_Error(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAccountById("82956508ccde2d63320a8a59")
	msg := "Erro ao listar Account Id"

	if err == nil {
		t.Error("Get Account by Id, Expected ERROR", ", Got Id:", a.Id)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Get Account by Id, Expected error message '", msg, "', Got:", err)
	}
}
