package auth

//
// microservices => auth => usecase => implementation_mock.go
//

import (
	repository "BackEnd_Api/microservices/auth/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var isSucesso bool

type mockAuthRepository struct {
	fakePassword string
}

func NewMockAuthRepository(fakePassword string) repository.IAuthRepository {
	return &mockAuthRepository{
		fakePassword: fakePassword,
	}
}

func (mIAR *mockAuthRepository) GetAccountPasswordHashByIdRepository(id string) (string, error) {

	if isSucesso {
		hahsPassord, err := bcrypt.GenerateFromPassword([]byte(mIAR.fakePassword), 14)

		if err != nil {
			return "", err
		}

		return string(hahsPassord), nil

	} else {
		return "", fmt.Errorf("Não existe a Account referida pelo Id : %s ", id)
	}

}

func (mIAR *mockAuthRepository) GetAccountPasswordHashByEmailRepository(email string) (string, error) {

	if isSucesso {
		hahsPassord, err := bcrypt.GenerateFromPassword([]byte(mIAR.fakePassword), 14)

		if err != nil {
			return "", err
		}

		return string(hahsPassord), nil

	} else {
		return "", fmt.Errorf("Não existe a Account referida pelo Email : %s ", email)
	}

}

func (mIAR *mockAuthRepository) GetAccountPasswordHashByDocumentRepository(document string) (string, error) {

	if isSucesso {
		hahsPassord, err := bcrypt.GenerateFromPassword([]byte(mIAR.fakePassword), 14)

		if err != nil {
			return "", err
		}

		return string(hahsPassord), nil

	} else {
		return "", fmt.Errorf("Não existe a Account referida pelo Document : %s ", document)
	}

}

func (mIAR *mockAuthRepository) GetAccountPasswordHashByPhoneRepository(phone string) (string, error) {

	if isSucesso {
		hahsPassord, err := bcrypt.GenerateFromPassword([]byte(mIAR.fakePassword), 14)

		if err != nil {
			return "", err
		}

		return string(hahsPassord), nil

	} else {
		return "", fmt.Errorf("Não existe a Account referida pelo Phone : %s ", phone)
	}

}
