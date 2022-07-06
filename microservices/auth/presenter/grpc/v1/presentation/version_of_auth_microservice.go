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

func (aS *AuthMicroService) VersionOfAuthMicroService(ctx context.Context, in *pb.VersionOfAuthMicroServiceRequest) (*pb.VersionOfAuthMicroServiceResponse, error) {

	if in == nil {
		return nil, status.Error(codes.FailedPrecondition, "Request não foi Fornecidos")
	}

	version := aS.Usecase.VersionOfAuthMicroService()

	r := pb.VersionOfAuthMicroServiceResponse{
		Version: version,
	}

	return &r, nil
}
