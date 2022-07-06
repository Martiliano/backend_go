package account

//
// microservices => account => external => db => mongo => implementation => get_account_by_email_repository.go
//

import (
	account_model "BackEnd_Api/microservices/account/external/db/mongo/model"
	account "BackEnd_Api/microservices/account/rules"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sMR *accountMongoRepository) GetAccountByEmailRepository(email string) (*account.Account, error) {
	collection := sMR.dbServer.Database(sMR.dbName).Collection(collectionName)

	result := account_model.AccountDataStruct{}
	a := account.Account{}

	filter := bson.D{primitive.E{Key: "email", Value: email}}

	err := collection.FindOne(sMR.ctx, filter).Decode(&result)

	if err != nil {
		return &a, err
	}

	a.Id = fmt.Sprintf("%s", result.Id)[10:34]
	a.UserName = result.UserName
	a.PhoneNumber = result.PhoneNumber
	a.Email = result.Email
	a.Document = result.Document
	a.FullName = result.FullName
	a.TypeOfAccount = result.TypeOfAccount
	a.Roles = result.Roles

	return &a, nil
}
