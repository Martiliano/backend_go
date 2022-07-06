package jwt

//
// microservices => helpers => jwt => jwt.go
//

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtManager struct {
	SecretKey     string
	TokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	TypeOfAccount []string `json:"type_of_account"`
	Roles         []string `json:"roles"`
}

func NewJwtManager(secretKey string, tokenDuration time.Duration) *JwtManager {
	return &JwtManager{secretKey, tokenDuration}
}

func (j *JwtManager) Generate(id string, typeOfAccount []string, roles []string) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        id,
			Issuer:    "backend_go.v1",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(j.TokenDuration).Unix(),
		},
		TypeOfAccount: typeOfAccount,
		Roles:         roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JwtManager) GenerateWithDetails(id string, typeOfAccount []string, roles []string, issuer string, issuedAt int64, expiresAt int64) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        id,
			Issuer:    issuer,
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		},
		TypeOfAccount: typeOfAccount,
		Roles:         roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JwtManager) GenerateSetPassword(id string, typeOfAccount []string, roles []string) (string, error) {

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        id,
			Issuer:    "backend_go.v1.setpassword",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(j.TokenDuration).Unix(),
		},
		TypeOfAccount: typeOfAccount,
		Roles:         roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JwtManager) GenerateRecoveryPassword(id string, typeOfAccount []string, roles []string) (string, error) {

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        id,
			Issuer:    "backend_go.v1.recoverypassword",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(j.TokenDuration).Unix(),
		},
		TypeOfAccount: typeOfAccount,
		Roles:         roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

func (j *JwtManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Metodo de assinatura desconhecido")
			}

			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Token invalido: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("Claims do Token são invalidos")
	}

	return claims, nil
}
