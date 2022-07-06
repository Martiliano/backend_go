package account

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	"BackEnd_Api/logger"
	"testing"

	"go.uber.org/zap"
)

//
// microservices => account => usecase => update_account_test.go
//

func TestVersionOfAccountMicroService_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	version := usecase.VersionOfAccountMicroService()
	expected := "0.1.0-Account"

	if version != expected {
		t.Error("Version Of Account MicroService, Expected '", expected, "' :", version)
	}
}
