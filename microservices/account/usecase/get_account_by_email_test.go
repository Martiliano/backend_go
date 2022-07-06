package account

//
// microservices => account => usecase => get_account_by_email_test.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	"BackEnd_Api/logger"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func TestGetAccountByEmail_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAccountByEmail("anyone@email.com.br")

	if err != nil {
		t.Error("Get Account by Document, Expected Account Id :", "69856696542", "Got :", err)
	} else if a.Id == "69856696542" {
		t.Error("Get Account by Document, Expected Account Id :", "69856696542", "Got :", a.Id)
	}
}

func TestGetAccountByEmail_Error(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAccountByEmail("anyone@email.com")
	msg := "Não existe a Account referida pelo email"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", a.Id)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}
