package account

//
// microservices => account => presenter => grpc => v1 => presentation => update_account.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	account "BackEnd_Api/microservices/account/rules"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AccountMicroService) UpdateAccount(ctx context.Context, in *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {

	if in == nil || in.Account == nil {
		return nil, status.Error(codes.FailedPrecondition, "Account não foi Fornecidos")
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

	count, err := aS.Usecase.UpdateAccount(a)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	r := pb.UpdateAccountResponse{
		Success:         true,
		NumberOfUpdates: int32(count),
		Msg:             "Account alterado com sucesso, Id: " + in.GetAccount().GetId(),
	}

	return &r, nil
}
