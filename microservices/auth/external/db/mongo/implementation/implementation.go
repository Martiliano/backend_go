package auth

import (
	auth_repository "BackEnd_Api/microservices/auth/repository"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//
// microservices => auth => external => db => mongo => implementation => implementation.go
//

type authMongoRepository struct {
	dbServer *mongo.Client
	ctx      context.Context
	dbName   string
}

const collectionAccountName string = "account"

func NewAuthMongoRepository(db *mongo.Client, ctx context.Context, dbName string) auth_repository.IAuthRepository {
	return &authMongoRepository{dbServer: db, ctx: ctx, dbName: dbName}
}
