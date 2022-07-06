package account

//
// microservices => account => presenter => grpc => v1 => presentation => get_all_accounts.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AccountMicroService) GetAllAccounts(ctx context.Context, in *pb.GetAllAccountsRequest) (*pb.GetAllAccountsResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Resquisição não foi informada")
	}

	a, err := aS.Usecase.GetAllAccounts()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ai := []*pb.AccountEntities{}

	for _, result := range *(a) {
		t := pb.AccountEntities{}

		t.Id = result.Id
		t.UserName = result.UserName
		t.PhoneNumber = result.PhoneNumber
		t.Email = result.Email
		t.Document = result.Document
		t.FullName = result.FullName
		t.TypeOfAccount = result.TypeOfAccount
		t.Roles = result.Roles

		ai = append(ai, &t)
	}

	r := pb.GetAllAccountsResponse{
		Success: true,
		Msg:     "Dados de todas as contas",
		Account: ai,
	}

	return &r, nil
}
