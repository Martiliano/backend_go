package account

//
// microservices => account => presenter => grpc => v1 => presentation => recovery_password.go
//
//

import (
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"

	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (aS *AccountMicroService) RecoveryPassword(ctx context.Context, in *pb.RecoveryPasswordRequest) (*pb.RecoveryPasswordResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Resquisição não foi informada")
	}

	err := aS.Usecase.RecoveryPassword(in.GetToken(), in.GetPassword(), in.GetConfirmPassword())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	r := pb.RecoveryPasswordResponse{
		Success: true,
		Msg:     "Token gerado com sucesso.",
	}

	return &r, nil
}
