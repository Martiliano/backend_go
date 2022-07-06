package account

//
// microservices => account => external => db => mongo => implementation => create_account_repository.go
//

import (
	account_model "BackEnd_Api/microservices/account/external/db/mongo/model"
	account "BackEnd_Api/microservices/account/rules"

	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sMR *accountMongoRepository) CreateAccountRepository(a account.Account) (string, error) {

	collection := sMR.dbServer.Database(sMR.dbName).Collection(collectionName)

	doc := account_model.AccountDataStructInsert{
		UserName:      a.UserName,
		PhoneNumber:   a.PhoneNumber,
		Email:         a.Email,
		Document:      a.Document,
		FullName:      a.FullName,
		TypeOfAccount: a.TypeOfAccount,
		Roles:         a.Roles,

		Active:  true,
		Blocked: false,
		Removed: false,

		LoginInOnlyDevice: false,
		AccessToken:       "",

		PasswordHash:     "",
		PasswordStatus:   "",
		PasswordValidity: "",
		ResetPasswordAt:  primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),

		TokenSetPassword: "",
		PasswordSetedAt:  primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),

		TokenPasswordRecovery: "",
		PasswordRecoveredAt:   primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),

		TokenResetPassword: "",
		ResetedPasswordAt:  primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),

		LastLoginAt: primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),

		ConnectedToDevice: false,
		ActiveDeviceId:    "",
		ActiveDeviceAt:    primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),

		RegisterOrigin: "",
		AvatarPath:     "",

		PhoneIsVerified:     false,
		ValidatedPhoneAt:    primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),
		TokenValidatedPhone: "",
		EmailIsVerified:     false,
		ValidatedEmailAt:    primitive.NewDateTimeFromTime(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)),
		TokenValidatedEmail: false,

		CreateAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdateAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	result, err := collection.InsertOne(sMR.ctx, doc)

	if err != nil {
		return "", err
	}

	strId := fmt.Sprintf("%s", result.InsertedID)[10:34]

	return strId, nil
}
