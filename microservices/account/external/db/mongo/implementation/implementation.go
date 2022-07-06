package account

//
// microservices => account => external => db => mongo => implementation => implementation.go
//

import (
	account_repository "BackEnd_Api/microservices/account/repository"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type accountMongoRepository struct {
	dbServer *mongo.Client
	ctx      context.Context
	dbName   string
}

const collectionName string = "account"

func NewAccountMongoRepository(db *mongo.Client, ctx context.Context, dbName string) account_repository.IAccountRepository {
	return &accountMongoRepository{dbServer: db, ctx: ctx, dbName: dbName}
}
