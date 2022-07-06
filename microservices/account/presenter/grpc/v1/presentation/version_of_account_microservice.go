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

func (aS *AccountMicroService) VersionOfAccountMicroService(ctx context.Context, in *pb.VersionOfAccountMicroServiceRequest) (*pb.VersionOfAccountMicroServiceResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request não foi Fornecidos")
	}

	version := aS.Usecase.VersionOfAccountMicroService()

	r := pb.VersionOfAccountMicroServiceResponse{
		Version: version,
	}
	return &r, nil
}
