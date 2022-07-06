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

func (aS *AuthMicroService) LoginByPhone(ctx context.Context, in *pb.LoginByPhoneRequest) (*pb.LoginByPhoneResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request não foi Fornecidos")
	}

	token, err := aS.Usecase.LoginByPhone(in.GetPhone(), in.GetPassword())

	if err != nil {
		return nil, err
	}

	r := pb.LoginByPhoneResponse{
		Success: true,
		Token:   token,
	}

	return &r, nil
}
