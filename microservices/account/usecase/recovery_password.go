package account

import (
	config "BackEnd_Api/config"
	jwt "BackEnd_Api/helpers/jwt"

	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//
// microservices => account => usecase => recovery_password.go
//

func (aUC *AccountUseCase) RecoveryPassword(token string, password string, confirmPassword string) error {

	if password != confirmPassword {
		return errors.New("A senha e a confirmação são diferentes")
	}

	cfg := config.GetConfig()
	jwtManager := jwt.NewJwtManager(cfg.Auth.CreateSecret, time.Hour*48)

	claims, err := jwtManager.Verify(token)

	if err != nil {
		return err
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return errors.New("Token Expirado")
	}

	if claims.Issuer != "backend_go.v1.recoverypassword" {
		return errors.New("Issuer Token Invalido")
	}

	account, err := aUC.repository.GetAccountByRecoveryTokenRepository(token)

	if err != nil {
		return err
	}

	hahsPassord, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	count, err := aUC.repository.RecoveryPasswordRepository(account.Id, string(hahsPassord))

	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("Usuário não encontrado")
	}

	return nil
}
