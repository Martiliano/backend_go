package auth

//
// microservices => auth => presenter => grpc => v1 => presentation => login_by_id.go
//
//

import (
	pb "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AuthMicroService) LoginById(ctx context.Context, in *pb.LoginByIdRequest) (*pb.LoginByIdResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request não foi Fornecidos")
	}

	token, err := aS.Usecase.LoginById(in.GetId(), in.GetPassword())

	if err != nil {
		return nil, err
	}

	r := pb.LoginByIdResponse{
		Success: true,
		Token:   token,
	}

	return &r, nil
}
