package auth

//
// microservices => auth => presenter => grpc => v1 => presentation => login_by_phone.go
//
//

import (
	pb "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AuthMicroService) LoginByDocument(ctx context.Context, in *pb.LoginByDocumentRequest) (*pb.LoginByDocumentResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request não foi Fornecidos")
	}

	token, err := aS.Usecase.LoginByDocument(in.GetDocument(), in.GetPassword())

	if err != nil {
		return nil, err
	}

	r := pb.LoginByDocumentResponse{
		Success: true,
		Token:   token,
	}

	return &r, nil
}
