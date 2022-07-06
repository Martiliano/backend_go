package auth

//
// microservices => auth => microservice.go
//

import (
	config "BackEnd_Api/config"
	account_repository_mongo "BackEnd_Api/microservices/account/external/db/mongo/implementation"
	auth_repository_mongo "BackEnd_Api/microservices/auth/external/db/mongo/implementation"
	pb "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"
	presentation "BackEnd_Api/microservices/auth/presenter/grpc/v1/presentation"
	auth_usecase "BackEnd_Api/microservices/auth/usecase"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthMicroService(config *config.Config, db *mongo.Client) pb.AuthServer {
	ctx := context.Background()

	auth_repository_mongo := auth_repository_mongo.NewAuthMongoRepository(db, ctx, config.Mongo.DbName)
	account_repository_mongo := account_repository_mongo.NewAccountMongoRepository(db, ctx, config.Mongo.DbName)

	auth_usecase := auth_usecase.NewAuthUseCase(auth_repository_mongo, account_repository_mongo, config)

	return presentation.NewAuthServer(auth_usecase)
}
