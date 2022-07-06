package auth

//
// microservices => account => usecase => login_by_document.go
//

import (
	config "BackEnd_Api/config"
	"BackEnd_Api/helpers/jwt"
	"errors"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (aUC *AuthUseCase) LoginByDocument(document string, password string) (string, error) {

	hash, err := aUC.auth_repository.GetAccountPasswordHashByDocumentRepository(document)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return "", err
	}

	a, err := aUC.account_repository.GetAccountByDocumentRepository(document)

	cfg := config.GetConfig()
	duration, err := strconv.Atoi(cfg.Auth.DurationInMinutes)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Login By Document, error: %s", err))
	}

	jwtManager := jwt.NewJwtManager(cfg.Auth.Secret, time.Hour*time.Duration(duration))

	if err != nil {
		return "", err
	}

	token, err := jwtManager.Generate(a.Id, a.TypeOfAccount, a.Roles)

	if err != nil {
		return "", err
	}

	return token, nil
}
