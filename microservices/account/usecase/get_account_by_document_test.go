package account

//
// microservices => account => usecase => get_account_by_document_test.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	"BackEnd_Api/logger"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func TestGetAccountByDocument_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAccountByDocument("69856696542")

	if err != nil {
		t.Error("Get Account by Document, Expected Account Id :", "69856696542", "Got :", err)
	} else if a.Id == "69856696542" {
		t.Error("Get Account by Document, Expected Account Id :", "69856696542", "Got :", a.Id)
	}
}

func TestGetAccountByDocument_Error(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	a, err := usecase.GetAccountByDocument("79856696542")
	msg := "Não existe a Account referida pelo documento"

	if err == nil {
		t.Error("Get Account by Document, Expected error message '", msg, "', Got:", a.Document)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Get Account by Document, Expected error message '", msg, "', Got:", err)
	}
}
