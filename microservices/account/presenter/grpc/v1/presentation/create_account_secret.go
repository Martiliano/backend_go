package account

//
// microservices => account => presenter => grpc => v1 => presentation => create_account_secret.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	account "BackEnd_Api/microservices/account/rules"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AccountMicroService) CreateAccountSecret(ctx context.Context, in *pb.CreateAccountSecretRequest) (*pb.CreateAccountSecretResponse, error) {

	if in == nil || in.Account == nil {
		return nil, status.Error(codes.FailedPrecondition, "Account ou Secret não foram Fornecidos")
	}

	a := account.Account{
		Id:            in.GetAccount().GetId(),
		UserName:      in.GetAccount().GetUserName(),
		PhoneNumber:   in.GetAccount().GetPhoneNumber(),
		Email:         in.GetAccount().GetEmail(),
		Document:      in.GetAccount().GetDocument(),
		FullName:      in.GetAccount().GetFullName(),
		TypeOfAccount: in.GetAccount().GetTypeOfAccount(),
		Roles:         in.GetAccount().GetRoles(),
	}

	id, err := aS.Usecase.CreateAccountBySecret(a, in.GetSecret())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	r := pb.CreateAccountSecretResponse{
		Success: true,
		NewId:   id,
		Msg:     "Account incluido com sucesso, Id: " + id,
	}

	return &r, nil
}
