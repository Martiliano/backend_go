package auth

//
// microservices => auth => external => db => mongo => implementation => get_account_password_hash_by_phone.go
//

import (
	account_model "BackEnd_Api/microservices/account/external/db/mongo/model"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sMR *authMongoRepository) GetAccountPasswordHashByPhoneRepository(phone string) (string, error) {
	collection := sMR.dbServer.Database(sMR.dbName).Collection(collectionAccountName)

	result := account_model.AccountDataStruct{}

	filter := bson.D{primitive.E{Key: "phone_number", Value: phone}}

	err := collection.FindOne(sMR.ctx, filter).Decode(&result)

	if err != nil {
		return "", err
	}

	if result.PasswordStatus != "Valid" {
		return "", errors.New("Password Hash Invalida")
	}

	return result.PasswordHash, nil
}
