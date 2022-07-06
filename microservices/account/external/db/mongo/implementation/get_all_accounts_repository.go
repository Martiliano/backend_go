package account

//
// microservices => account => external => db => mongo => implementation => get_all_accounts_repository.go
//

import (
	account_model "BackEnd_Api/microservices/account/external/db/mongo/model"
	account "BackEnd_Api/microservices/account/rules"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (sMR *accountMongoRepository) GetAllAccountsRepository() (*[]account.Account, error) {
	collection := sMR.dbServer.Database(sMR.dbName).Collection(collectionName)

	filter := bson.D{{}}

	a := []account.Account{}

	results, err := collection.Find(sMR.ctx, filter)

	if err != nil {
		return &a, err
	}

	for results.Next(sMR.ctx) {
		result := account_model.AccountDataStruct{}
		t := account.Account{}

		err := results.Decode(&result)

		if err != nil {
			return &a, err
		}

		t.Id = fmt.Sprintf("%s", result.Id)[10:34]
		t.UserName = result.UserName
		t.PhoneNumber = result.PhoneNumber
		t.Email = result.Email
		t.Document = result.Document
		t.FullName = result.FullName
		t.TypeOfAccount = result.TypeOfAccount
		t.Roles = result.Roles

		a = append(a, t)
	}

	results.Close(sMR.ctx)

	if len(a) == 0 {
		return &a, mongo.ErrNoDocuments
	}

	return &a, nil
}
