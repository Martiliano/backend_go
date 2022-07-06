package auth

//
// microservices => auth => presenter => grpc => v1 => presentation => presentation.go
//
//

import (
	pb "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"
	auth "BackEnd_Api/microservices/auth/rules"
)

type AuthMicroService struct {
	Usecase auth.IAuth
}

func NewAuthServer(usecase auth.IAuth) pb.AuthServer {
	return &AuthMicroService{Usecase: usecase}
}
