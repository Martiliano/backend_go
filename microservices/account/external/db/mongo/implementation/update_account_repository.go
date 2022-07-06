package account

//
// microservices => account => external => db => mongo => implementation => update_account_repository.go
//

import (
	account "BackEnd_Api/microservices/account/rules"

	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sMR *accountMongoRepository) UpdateAccountRepository(a account.Account) (int, error) {
	collection := sMR.dbServer.Database(sMR.dbName).Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(a.Id)

	if err != nil {
		return 0, err
	}

	fmt.Println("UpdateAccountRepository objID : ", objID)

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "user_name", Value: a.UserName},
		primitive.E{Key: "phone_number", Value: a.PhoneNumber},
		primitive.E{Key: "email", Value: a.Email},
		primitive.E{Key: "document", Value: a.Document},
		primitive.E{Key: "type_of_account", Value: a.TypeOfAccount},
		primitive.E{Key: "roles", Value: a.Roles},
		primitive.E{Key: "full_name", Value: a.FullName},
		primitive.E{Key: "update_at", Value: primitive.NewDateTimeFromTime(time.Now())},
	}}}

	result, err := collection.UpdateOne(sMR.ctx, filter, updater)

	if err != nil {
		return 0, err
	}

	if result.ModifiedCount == 0 {
		return 0, errors.New("Registro a ser alterado não encontrado")
	}

	return int(result.ModifiedCount), nil
}
