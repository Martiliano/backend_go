package account

//
// microservices => account => usecase => create_account_by_secret_test.go
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

func TestCreateAccountBySecret_Success(t *testing.T) {

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

	code, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")

	if err != nil {
		t.Error("Create Account By Secret, Expected new Account Id :", a.Id, "Got :", err)
	} else if code != a.Id {
		t.Error("Create Account By Secret, Expected new Account Id :", a.Id, "Got :", code)
	}
}

func TestCreateAccountBySecret_InvalidSecret(t *testing.T) {

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
		Document:      "00000000000",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	isSucesso = true

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "Palavra secreta Invalida"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccountBySecret_InvalidPhoneNumber(t *testing.T) {

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
		PhoneNumber:   "54988s76655",
		Email:         "anyone@email.com",
		Document:      "00000000000",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	isSucesso = true

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "Telefone Formato Invalido"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccountBySecret_InvalidEmail(t *testing.T) {

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
		Document:      "00000000000",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	isSucesso = true

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "Email com Formato Invalido"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccountBySecret_InvalidDocument(t *testing.T) {

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
		Document:      "00000000000",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	isSucesso = true

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "CPF ou CNPJ Invalido"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccountBySecret_InvalidRoles(t *testing.T) {

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

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "Roles/Permissões Invalidos"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccountBySecret_InvalidTypeOfAccounts(t *testing.T) {

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

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "Tipo de Contas Invalidos"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccountBySecret_ErrorEmailExists(t *testing.T) {

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

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "Erro já existe um usuário utilizando este e-mail"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccountBySecret_ErrorPhoneExists(t *testing.T) {

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

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "Erro já existe um usuário utilizando este telefone"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}

func TestCreateAccountBySecret_ErrorDocumentExists(t *testing.T) {

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

	_, err := usecase.CreateAccountBySecret(a, "sdkfjsdfkfsdfjks")
	msg := "Erro já existe um usuário utilizando este documento"

	if err == nil {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("Create Account By Secret, using invalid Secret, Expected error message '", msg, "', Got:", err)
	}
}
