package account

//
// microservices => account => usecase => generate_set_password_token_test.go
//

import (
	config "BackEnd_Api/config"
	"BackEnd_Api/logger"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func TestGenerateSetPasswordToken_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	message := make(chan string)
	quit := make(chan bool)
	var err error

	go func() {

		repo := NewMockAccountRepository("Token")
		email := NewEmailMockService(message)
		config := config.GetConfig()

		usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

		isSucesso = true

		err = usecase.GenerateSetPasswordToken("62956508ccde2d63320a8a59")

		<-quit

	}()

	result := <-message
	msg := "<h1>CADASTRO DE SENHA</h1>"

	quit <- true

	if err != nil {
		t.Error("Generate Account Password Token, Expected Error nil :", "Got :", err)
	} else if !strings.HasPrefix(result, msg) {
		t.Error("Generate Account Password Token, Expected Token :", "Got :", result)
	}
}

func TestGenerateSetPasswordToken_Error(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	message := make(chan string)
	var err error

	repo := NewMockAccountRepository("Token")
	email := NewEmailMockService(message)
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	err = usecase.GenerateSetPasswordToken("82956508ccde2d63320a8a59")
	msg := "Erro ao listar Account Id"

	if err == nil {
		t.Error("Generate Account Password Token, Expected Error 'Erro ao listar Account Id' :", "Got :", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Generate Account Password Token, Expected Error 'Erro ao listar Account Id' :", "Got :", err)
	}
}
