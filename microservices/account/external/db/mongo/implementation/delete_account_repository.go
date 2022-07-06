package account

//
// microservices => account => external => db => mongo => implementation => delete_account_repository.go
//

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sMR *accountMongoRepository) DeleteAccountRepository(id string) (int, error) {
	collection := sMR.dbServer.Database(sMR.dbName).Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return 0, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return 0, err
	}

	if result.DeletedCount == 0 {
		return 0, errors.New("Registro a ser excluido não encontrado")
	}

	return int(result.DeletedCount), nil
}
