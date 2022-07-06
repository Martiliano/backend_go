package account

//
// microservices => account => presenter => grpc => v1 => presentation => generate_set_password_token.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AccountMicroService) GenerateSetPasswordToken(ctx context.Context, in *pb.GenerateSetPasswordTokenRequest) (*pb.GenerateSetPasswordTokenResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Resquisição não foi informada")
	}

	err := aS.Usecase.GenerateSetPasswordToken(in.GetId())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	r := pb.GenerateSetPasswordTokenResponse{
		Success: true,
		Msg:     "Token de recuperação de senha gerado para o, Id: " + in.Id,
	}

	return &r, nil
}
