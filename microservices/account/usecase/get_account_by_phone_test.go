package account

//
// microservices => account => usecase => get_account_by_phone_test.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	"BackEnd_Api/logger"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func TestGetAccountByPhone_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAccountByPhone("54988776543")

	if err != nil {
		t.Error("Get Account by PhoneNumber, Expected Account PhoneNumber :", "54988776543", "Got :", err)
	} else if a.PhoneNumber != "54988776543" {
		t.Error("Get Account by PhoneNumber, Expected Account PhoneNumber :", "54988776543", "Got :", a.PhoneNumber)
	}
}

func TestGetAccountByPhone_Error(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAccountByPhone("55988776543")
	msg := "Não existe a Account referida pelo telefone"

	if err == nil {
		t.Error("Get Account by PhoneNumber, Expected ERROR", ", Got PhoneNumber:", a.PhoneNumber)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Get Account by PhoneNumber, Expected error message '", msg, "', Got:", err)
	}
}
