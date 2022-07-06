package account

//
// microservices => account => external => db => mongo => implementation => get_set_password_token.go
//

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sMR *accountMongoRepository) GenerateSetPasswordTokenRepository(id string, token string) error {
	collection := sMR.dbServer.Database(sMR.dbName).Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "token_set_password", Value: token},
		primitive.E{Key: "password_seted_at", Value: primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC))},
		primitive.E{Key: "update_at", Value: primitive.NewDateTimeFromTime(time.Now())},
	}}}

	result, err := collection.UpdateOne(sMR.ctx, filter, updater)

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("Registro a ser alterado não encontrado")
	}

	return nil
}
