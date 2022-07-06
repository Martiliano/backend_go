package grpc

//
// presenter => grpc => monolithic => register.go
//

import (
	"BackEnd_Api/config"
	accountMicroservice "BackEnd_Api/microservices/account"
	accountMicroservicePb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	authMicroservice "BackEnd_Api/microservices/auth"
	authMicroservicePb "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func RegisterGrpcModules(config *config.Config, srv *grpc.Server, db *mongo.Client) {

	accountMicroservicePb.RegisterAccountServer(srv, accountMicroservice.NewAccountMicroService(config, db))

	authMicroservicePb.RegisterAuthServer(srv, authMicroservice.NewAuthMicroService(config, db))
}
