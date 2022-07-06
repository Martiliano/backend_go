package account

//
// microservices => account => external => db => mongo => implementation => create_account_repository_test.go
//

import (
	account "BackEnd_Api/microservices/account/rules"

	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateAccountRepository_Success(t *testing.T) {
	ctx := context.Background()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		t.Error("Create Account, Expected new Account Id : - , Got :", err)
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		t.Error("Create Account, Expected new Account Id : - , Got :", err)
	}

	a := account.Account{
		UserName:      "anyone",
		PhoneNumber:   "54988776655",
		Email:         "anyone@email.com",
		Document:      "12345678900",
		FullName:      "Anyone Unknown",
		TypeOfAccount: []string{"Desenvolvedor"},
		Roles:         []string{"Desevolvedor", "Administrador"},
	}

	repository := NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")

	id, err := repository.CreateAccountRepository(a)

	if err == nil {
		objID, err2 := primitive.ObjectIDFromHex(id)

		if err2 == nil {
			collection := db.Database("BackEnd_Api_test").Collection(collectionName)

			filter := bson.D{primitive.E{Key: "_id", Value: objID}}

			collection.DeleteOne(context.TODO(), filter)
		}

	}

	if err != nil {
		t.Error("Create Account, Expected new Account Id :", id, "Got :", err)
	}
}
