package account

//
// microservices => account => usecase => get_set_password_token.go
//

import (
	jwt "BackEnd_Api/helpers/jwt"
	"fmt"
	"time"
)

func (aUC *AccountUseCase) GenerateSetPasswordToken(id string) error {

	jwtManager := jwt.NewJwtManager(aUC.config.Auth.CreateSecret, time.Hour*48)

	account, err := aUC.GetAccountById(id)

	if err != nil {
		return err
	}

	token, err := jwtManager.GenerateSetPassword(id, account.TypeOfAccount, account.Roles)

	if err != nil {
		return err
	}

	err = aUC.repository.GenerateSetPasswordTokenRepository(id, token)

	if err != nil {
		return err
	}

	body := fmt.Sprintf("<h1>CADASTRO DE SENHA</h1><p></p><p>Abaixo segue o TOKEN necessário para você cadastrar a sua senha no sistema.</p><p></p><h2>Token:</h2><p>%s</p>", token)

	aUC.email.SendEmailTo(account.Email, account.FullName)
	aUC.email.SendEmailSetBody("text/html", body)

	err = aUC.email.SendEmail()

	if err != nil {
		return err
	}

	return nil
}
