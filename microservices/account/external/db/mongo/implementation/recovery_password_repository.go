package account

//
// microservices => account => external => db => mongo => implementation => recovery_password_account_repository.go
//

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sMR *accountMongoRepository) RecoveryPasswordRepository(id string, passwordHash string) (int, error) {
	collection := sMR.dbServer.Database(sMR.dbName).Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "password_hash", Value: passwordHash},
		primitive.E{Key: "password_status", Value: "Valid"},
		primitive.E{Key: "password_validity", Value: "Undefined"},
		primitive.E{Key: "token_password_recovery", Value: ""},
		primitive.E{Key: "password_recovered_at", Value: primitive.NewDateTimeFromTime(time.Now())},
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
