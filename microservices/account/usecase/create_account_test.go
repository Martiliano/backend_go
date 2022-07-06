package account

//
// microservices => account => usecase => create_account_test.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	"BackEnd_Api/logger"
	account "BackEnd_Api/microservices/account/rules"
	"strings"
	"testing"

	"go.uber.org/zap"
)

func TestCreateAccount_Success(t *testing.T) {

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

	code, err := usecase.CreateAccount(a)

	if err != nil {
		t.Error("Create Account, Expected new Account Id :", a.Id, "Got :", err)
	} else if code != a.Id {
		t.Error("Create Account, Expected new Account Id :", a.Id, "Got :", code)
	}
}

func TestCreateAccount_InvalidPhoneNumber(t *testing.T) {

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

	_, err := usecase.CreateAccount(a)
	msg := "Telefone Formato Invalido"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccount_InvalidEmail(t *testing.T) {

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

	_, err := usecase.CreateAccount(a)
	msg := "Email com Formato Invalido"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccount_InvalidDocument(t *testing.T) {

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
		Document:      "63849218004",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.CreateAccount(a)
	msg := "CPF ou CNPJ Invalido"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccount_InvalidRoles(t *testing.T) {

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

	_, err := usecase.CreateAccount(a)
	msg := "Roles/Permissões Invalidos"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccount_InvalidTypeOfAccounts(t *testing.T) {

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

	_, err := usecase.CreateAccount(a)
	msg := "Tipo de Contas Invalidos"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccount_ErrorEmailExists(t *testing.T) {

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
		Email:         "anyone@email.com.br",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.CreateAccount(a)
	msg := "Erro já existe um usuário utilizando este e-mail"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccount_ErrorPhoneExists(t *testing.T) {

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
		PhoneNumber:   "54988776543",
		Email:         "anyone@email.com",
		Document:      "04060664421",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.CreateAccount(a)
	msg := "Erro já existe um usuário utilizando este telefone"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccount_ErrorDocumentExists(t *testing.T) {

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
		Document:      "69856696542",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	isSucesso = true

	_, err := usecase.CreateAccount(a)
	msg := "Erro já existe um usuário utilizando este documento"

	if err == nil {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Create Account, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}
