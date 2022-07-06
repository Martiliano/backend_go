package auth

//
// microservices => auth => presenter => grpc => v1 => presentation => login_by_email.go
//
//

import (
	pb "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AuthMicroService) LoginByEmail(ctx context.Context, in *pb.LoginByEmailRequest) (*pb.LoginByEmailResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request não foi Fornecidos")
	}

	token, err := aS.Usecase.LoginByEmail(in.GetEmail(), in.GetPassword())

	if err != nil {
		return nil, err
	}

	r := pb.LoginByEmailResponse{
		Success: true,
		Token:   token,
	}

	return &r, nil
}
