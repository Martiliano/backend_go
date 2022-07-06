package account

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	jwt "BackEnd_Api/helpers/jwt"
	"BackEnd_Api/logger"
	"strings"
	"testing"
	"time"

	"go.uber.org/zap"
)

//
// microservices => account => usecase => recovery_password_test.go
//

func TestRecoveryPassword_PasswordNotMatch(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	jwtManager := jwt.NewJwtManager(config.Auth.CreateSecret, time.Hour*48)

	token, err := jwtManager.GenerateWithDetails("62956508ccde2d63320a8a59", []string{"Desenvolvedor"},
		[]string{"Desevolvedor", "Administrador"}, "backend_go.v1.recoverypassword", time.Now().Unix(),
		time.Now().Add(time.Millisecond*200).Unix())

	password := "strongPassword"
	confirmPassword := "strongPasswordError"

	err = usecase.RecoveryPassword(token, password, confirmPassword)
	msg := "A senha e a confirmação são diferentes"

	if err == nil {
		t.Error("RecoveryPassword, Expected error message: '", msg, "', Got:", err)
	} else if err.Error() != msg {
		t.Error("RecoveryPassword, Expected error message: '", msg, "', Got:", err)
	}
}

func TestRecoveryPassword_TokenExpired(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	jwtManager := jwt.NewJwtManager(config.Auth.CreateSecret, time.Hour*48)

	token, err := jwtManager.GenerateWithDetails("62956508ccde2d63320a8a59", []string{"Desenvolvedor"},
		[]string{"Desevolvedor", "Administrador"}, "backend_go.v1.recoverypassword", time.Now().Unix(),
		time.Now().Add(time.Millisecond*300).Unix())

	time.Sleep(800 * time.Millisecond)

	password := "strongPassword"
	confirmPassword := "strongPassword"

	err = usecase.RecoveryPassword(token, password, confirmPassword)
	msg := "Token invalido: token is expired by"

	if err == nil {
		t.Error("RecoveryPassword, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("RecoveryPassword, Expected error message '", msg, "', Got:", err)
	}
}

func TestRecoveryPassword_InvalidIssuer(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	repo := NewMockAccountRepository("Token")
	email := email_service.NewEmailService()
	config := config.GetConfig()

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	jwtManager := jwt.NewJwtManager(config.Auth.CreateSecret, time.Hour*48)

	token, err := jwtManager.GenerateWithDetails("62956508ccde2d63320a8a59", []string{"Desenvolvedor"},
		[]string{"Desevolvedor", "Administrador"}, "iuris.invalid.issuer", time.Now().Unix(),
		time.Now().Add(time.Second*45).Unix())

	password := "strongPassword"
	confirmPassword := "strongPassword"

	err = usecase.RecoveryPassword(token, password, confirmPassword)
	msg := "Issuer Token Invalido"

	if err == nil {
		t.Error("RecoveryPassword, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("RecoveryPassword, Expected error message '", msg, "', Got:", err)
	}
}

func TestRecoveryPassword_TokenNotFound(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	email := email_service.NewEmailService()
	config := config.GetConfig()

	jwtManager := jwt.NewJwtManager(config.Auth.CreateSecret, time.Hour*48)

	token, err := jwtManager.GenerateWithDetails("62956508ccde2d63320a8a59", []string{"Desenvolvedor"},
		[]string{"Desevolvedor", "Administrador"}, "backend_go.v1.recoverypassword", time.Now().Unix(),
		time.Now().Add(time.Second*45).Unix())

	repo := NewMockAccountRepository("Token")

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	password := "strongPassword"
	confirmPassword := "strongPassword"

	err = usecase.RecoveryPassword(token, password, confirmPassword)
	msg := "Não existe a Account referida pelo token"

	if err == nil {
		t.Error("RecoveryPassword, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("RecoveryPassword, Expected error message '", msg, "', Got:", err)
	}
}

func TestRecoveryPassword_ErrorUpdatePassword(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	email := email_service.NewEmailService()
	config := config.GetConfig()

	jwtManager := jwt.NewJwtManager(config.Auth.CreateSecret, time.Hour*48)

	token, err := jwtManager.GenerateWithDetails("62956508ccde2d63320a8a59", []string{"Desenvolvedor"},
		[]string{"Desevolvedor", "Administrador"}, "backend_go.v1.recoverypassword", time.Now().Unix(),
		time.Now().Add(time.Second*45).Unix())

	repo := NewMockAccountRepository(token)

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = false

	password := "strongPassword"
	confirmPassword := "strongPassword"

	err = usecase.RecoveryPassword(token, password, confirmPassword)
	msg := "Não foi possivel Recuperar a senha para o Id"

	if err == nil {
		t.Error("RecoveryPassword, Expected error message '", msg, "', Got:", err)
	} else if !strings.HasPrefix(err.Error(), msg) {
		t.Error("RecoveryPassword, Expected error message '", msg, "', Got:", err)
	}
}

func TestRecoveryPassword_Success(t *testing.T) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
		Development: false,
	})

	email := email_service.NewEmailService()
	config := config.GetConfig()

	jwtManager := jwt.NewJwtManager(config.Auth.CreateSecret, time.Hour*48)

	token, err := jwtManager.GenerateWithDetails("62956508ccde2d63320a8a59", []string{"Desenvolvedor"},
		[]string{"Desevolvedor", "Administrador"}, "backend_go.v1.recoverypassword", time.Now().Unix(),
		time.Now().Add(time.Second*45).Unix())

	repo := NewMockAccountRepository(token)

	usecase := NewAccountUseCase(repo, email, config, "sdkfjsdfkfsdfjks")

	isSucesso = true

	password := "strongPassword"
	confirmPassword := "strongPassword"

	err = usecase.RecoveryPassword(token, password, confirmPassword)

	if err != nil {
		t.Error("RecoveryPassword, Expected SUCCESS, Got:", err)
	}
}
