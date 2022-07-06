package account

//
// microservices => account => presenter => grpc => v1 => presentation => generate_recovery_password_token.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AccountMicroService) GenerateRecoveryPasswordToken(ctx context.Context, in *pb.GenerateRecoveryPasswordTokenRequest) (*pb.GenerateRecoveryPasswordTokenResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Resquisição não foi informada")
	}

	err := aS.Usecase.GenerateRecoveryPasswordToken(in.GetId())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	r := pb.GenerateRecoveryPasswordTokenResponse{
		Success: true,
		Msg:     "Token de recuperação de senha gerado para o, Id: " + in.Id,
	}

	return &r, nil
}
