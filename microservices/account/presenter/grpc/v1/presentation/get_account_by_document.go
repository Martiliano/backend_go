package account

//
// microservices => account => presenter => grpc => v1 => presentation => get_account_by_document.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AccountMicroService) GetAccountByDocument(ctx context.Context, in *pb.GetAccountByDocumentRequest) (*pb.GetAccountByDocumentResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Resquisição não foi informada")
	}

	a, err := aS.Usecase.GetAccountByDocument(in.GetDocument())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ai := &pb.AccountEntities{
		Id:            a.Id,
		UserName:      a.UserName,
		PhoneNumber:   a.PhoneNumber,
		Email:         a.Email,
		Document:      a.Document,
		FullName:      a.FullName,
		TypeOfAccount: a.TypeOfAccount,
		Roles:         a.Roles,
	}

	r := pb.GetAccountByDocumentResponse{
		Success: true,
		Msg:     "Dados da conta do Id: " + a.Id,
		Account: ai,
	}

	return &r, nil
}
