package account

//
// microservices => account => usecase => implementation_mock.go
//

import (
	repository "BackEnd_Api/microservices/account/repository"
	account "BackEnd_Api/microservices/account/rules"
	"fmt"
)

var isSucesso bool

type mockAccountRepository struct {
	token string
}

func NewMockAccountRepository(token string) repository.IAccountRepository {
	return &mockAccountRepository{
		token: token,
	}
}

func (mIAR *mockAccountRepository) CreateAccountRepository(a account.Account) (string, error) {

	if isSucesso {
		return a.Id, nil
	} else {
		return "", fmt.Errorf("Erro ao criar Account : %s ", a.FullName)
	}

}

func (mIAR *mockAccountRepository) GetAccountByIdRepository(id string) (*account.Account, error) {

	a := account.Account{
		Id:            "62956508ccde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776543",
		Email:         "anyone@email.com.br",
		Document:      "69856696542",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	if isSucesso {
		if id == a.Id {
			return &a, nil
		} else {
			return nil, fmt.Errorf("Erro ao listar Account Id : %s ", id)
		}

	} else {
		return nil, fmt.Errorf("Erro ao listar Account Id : %s ", id)
	}

}

func (mIAR *mockAccountRepository) GetAllAccountsRepository() (*[]account.Account, error) {

	all := []account.Account{
		{
			Id:            "62956508ccde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776543",
			Email:         "anyone@email.com.br",
			Document:      "69856696542",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		{
			Id:            "62956508ccde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776543",
			Email:         "anyone@email.com.br",
			Document:      "69856696542",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
		{
			Id:            "62956508ccde2d63320a8a59",
			UserName:      "anyone",
			PhoneNumber:   "54988776543",
			Email:         "anyone@email.com.br",
			Document:      "69856696542",
			FullName:      "Anyone Unknown",
			TypeOfAccount: []string{"Desenvolvedor"},
			Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
		},
	}

	if isSucesso {
		return &all, nil
	} else {
		return nil, fmt.Errorf("Erro ao listar todas as Account.")
	}
}

func (mIAR *mockAccountRepository) UpdateAccountRepository(a account.Account) (int, error) {

	if isSucesso {
		return 1, nil
	} else {
		return 0, fmt.Errorf("Erro ao Atualizar Account : %s ", a.FullName)
	}

}

func (mIAR *mockAccountRepository) DeleteAccountRepository(id string) (int, error) {

	if isSucesso {
		return 1, nil
	} else {
		return 0, fmt.Errorf("Erro ao Deletar Account Id : %s ", id)
	}

}

func (mIAR *mockAccountRepository) GetAccountByEmailRepository(email string) (*account.Account, error) {

	a := account.Account{
		Id:            "62956508ccde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776543",
		Email:         "anyone@email.com.br",
		Document:      "69856696542",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	if isSucesso {
		if email == a.Email {
			return &a, nil
		} else {
			return nil, fmt.Errorf("Não existe a Account referida pelo email : %s ", email)
		}
	} else {
		return nil, fmt.Errorf("Não existe a Account referida pelo email : %s ", email)
	}

}

func (mIAR *mockAccountRepository) GetAccountByPhoneRepository(phone string) (*account.Account, error) {

	a := account.Account{
		Id:            "62956508ccde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776543",
		Email:         "anyone@email.com.br",
		Document:      "69856696542",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	if isSucesso {
		if phone == a.PhoneNumber {
			return &a, nil
		} else {
			return nil, fmt.Errorf("Não existe a Account referida pelo telefone : %s ", phone)
		}
	} else {
		return nil, fmt.Errorf("Não existe a Account referida pelo telefone : %s ", phone)
	}

}

func (mIAR *mockAccountRepository) GetAccountByDocumentRepository(document string) (*account.Account, error) {

	a := account.Account{
		Id:            "62956508ccde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776543",
		Email:         "anyone@email.com.br",
		Document:      "69856696542",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	if isSucesso {
		if document == a.Document {
			return &a, nil
		} else {
			return nil, fmt.Errorf("Não existe a Account referida pelo documento : %s ", document)
		}
	} else {
		return nil, fmt.Errorf("Não existe a Account referida pelo documento : %s ", document)
	}

}

func (mIAR *mockAccountRepository) GetAccountBySetPasswordTokenRepository(token string) (*account.Account, error) {

	a := account.Account{
		Id:            "62956508ccde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776543",
		Email:         "anyone@email.com.br",
		Document:      "69856696542",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	// isSucesso
	if token == mIAR.token {
		return &a, nil
	} else {
		return nil, fmt.Errorf("Não existe a Account referida pelo token : %s ", token)
	}

}

func (mIAR *mockAccountRepository) GetAccountByRecoveryTokenRepository(token string) (*account.Account, error) {

	a := account.Account{
		Id:            "62956508ccde2d63320a8a59",
		UserName:      "anyone",
		PhoneNumber:   "54988776543",
		Email:         "anyone@email.com.br",
		Document:      "69856696542",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"SemRestricoes", "Incluir", "Excluir", "Alterar", "Consultar", "Visualizar"},
	}

	// isSucesso
	if token == mIAR.token {
		return &a, nil
	} else {
		return nil, fmt.Errorf("Não existe a Account referida pelo token : %s ", token)
	}

}

func (mIAR *mockAccountRepository) SetPasswordRepository(id string, passwordHash string) (int, error) {

	if isSucesso {
		return 1, nil
	} else {
		return 0, fmt.Errorf("Não foi possivel setar a senha para o Id : %s ", id)
	}

}

func (mIAR *mockAccountRepository) GenerateSetPasswordTokenRepository(id string, token string) error {

	if isSucesso {
		return nil
	} else {
		return fmt.Errorf("Não foi possivel recuperar o token para o Id : %s ", id)
	}

}

func (mIAR *mockAccountRepository) RecoveryPasswordRepository(id string, passwordHash string) (int, error) {

	if isSucesso {
		return 1, nil
	} else {
		return 0, fmt.Errorf("Não foi possivel Recuperar a senha para o Id : %s ", id)
	}

}

func (mIAR *mockAccountRepository) GenerateRecoveryPasswordTokenRepository(id string, token string) error {

	if isSucesso {
		return nil
	} else {
		return fmt.Errorf("Não foi possivel recuperar o token para o Id : %s ", id)
	}

}
