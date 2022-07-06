package account

//
// microservices => account => usecase => update_account_test.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	"BackEnd_Api/logger"
	account "BackEnd_Api/microservices/account/rules"
	"testing"

	"go.uber.org/zap"
)

func TestUpdateAccount_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
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
		t.Error("Create Account, Expected 1 register updated Got :", err)
	} else if count != 1 {
		t.Error("Create Account, Expected 1 register updated Got :", count)
	}
}

func TestUpdateAccount_InvalidPhoneNumber(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988x76655",
		Email:         "anyone@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.UpdateAccount(a)
	msg := "Telefone Formato Invalido"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestUpdateAccount_InvalidEmail(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776655",
		Email:         "anyone@emailcom",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.UpdateAccount(a)
	msg := "Email com Formato Invalido"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestUpdateAccount_InvalidDocument(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776655",
		Email:         "anyone@email.com",
		Document:      "63849218224",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.UpdateAccount(a)
	msg := "CPF ou CNPJ Invalido"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestUpdateAccount_InvalidRoles(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
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
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Blablabla", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.UpdateAccount(a)
	msg := "Roles/Permissões Invalidos"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestUpdateAccount_InvalidTypeOfAccounts(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	a := account.Account{
		Id:            "62956508bbde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776655",
		Email:         "anyone@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor", "Blabla"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.UpdateAccount(a)
	msg := "Tipo de Contas Invalidos"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}
