package account

//
// microservices => account => presenter => grpc => v1 => presentation => delete_account.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AccountMicroService) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Resquisição não foi informada")
	}

	count, err := aS.Usecase.DeleteAccount(in.Id)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	r := pb.DeleteAccountResponse{
		Success:         true,
		NumberOfDeletes: int32(count),
		Msg:             "Account excluido com sucesso, Id: " + in.Id,
	}

	return &r, nil
}
