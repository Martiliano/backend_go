package account

//
// microservices => account => usecase => get_recovery_password_token.go
//

import (
	jwt "BackEnd_Api/helpers/jwt"
	"fmt"
	"time"
)

func (aUC *AccountUseCase) GenerateRecoveryPasswordToken(id string) error {

	jwtManager := jwt.NewJwtManager(aUC.config.Auth.CreateSecret, time.Hour*48)

	account, err := aUC.GetAccountById(id)

	if err != nil {
		return err
	}

	token, err := jwtManager.GenerateRecoveryPassword(id, account.TypeOfAccount, account.Roles)

	if err != nil {
		return err
	}

	err = aUC.repository.GenerateRecoveryPasswordTokenRepository(id, token)

	if err != nil {
		return err
	}

	body := fmt.Sprintf("<h1>RECUPERAÇÃO DE SENHA</h1><p></p><p>Abaixo segue o TOKEN necessário para você recuperar a sua senha no sistema.</p><p></p><h2>Token:</h2><p>%s</p>", token)

	aUC.email.SendEmailTo(account.Email, account.FullName)
	aUC.email.SendEmailSetBody("text/html", body)

	err = aUC.email.SendEmail()

	if err != nil {
		return err
	}

	return nil
}
