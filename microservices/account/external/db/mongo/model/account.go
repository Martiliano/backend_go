package account_model

//
// microservices => account => external => db => mongo => model => implementation.go
//

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountDataStruct struct {
	Id            primitive.ObjectID `bson:"_id"`
	UserName      string             `bson:"user_name"`
	PhoneNumber   string             `bson:"phone_number"`
	Email         string             `bson:"email"`
	Document      string             `bson:"document"`
	FullName      string             `bson:"full_name"`
	TypeOfAccount []string           `bson:"type_of_account"`
	Roles         []string           `bson:"roles"`

	Active  bool `bson:"active"`
	Blocked bool `bson:"blocked"`
	Removed bool `bson:"removed"`

	LoginInOnlyDevice bool   `bson:"login_in_only_device"`
	AccessToken       string `bson:"access_token"`

	PasswordHash     string             `bson:"password_hash"`
	PasswordStatus   string             `bson:"password_status"`
	PasswordValidity string             `bson:"password_validity"`
	ResetPasswordAt  primitive.DateTime `bson:"reset_password_at"`

	TokenSetPassword string             `bson:"token_set_password"`
	PasswordSetedAt  primitive.DateTime `bson:"password_seted_at"`

	TokenPasswordRecovery string             `bson:"token_password_recovery"`
	PasswordRecoveredAt   primitive.DateTime `bson:"password_recovered_at"`

	TokenResetPassword string             `bson:"token_reset_password"`
	ResetedPasswordAt  primitive.DateTime `bson:"reseted_password_at"`

	LastLoginAt primitive.DateTime `bson:"last_login_at"`

	ConnectedToDevice bool               `bson:"connected_to_device"`
	ActiveDeviceId    string             `bson:"ActiveDeviceId"`
	ActiveDeviceAt    primitive.DateTime `bson:"activeDeviceAt"`

	RegisterOrigin string `bson:"register_origin"`
	AvatarPath     string `bson:"avatar_path"`

	PhoneIsVerified     bool               `bson:"phone_is_verified"`
	ValidatedPhoneAt    primitive.DateTime `bson:"validated_phone_at"`
	TokenValidatedPhone string             `bson:"token_validated_phone"`

	EmailIsVerified     bool               `bson:"email_is_verified"`
	ValidatedEmailAt    primitive.DateTime `bson:"validated_email_at"`
	TokenValidatedEmail bool               `bson:"token_validated_email"`

	CreateAt primitive.DateTime `bson:"create_at"`
	UpdateAt primitive.DateTime `bson:"update_at"`
}

type AccountDataStructInsert struct {
	UserName      string   `bson:"user_name"`
	PhoneNumber   string   `bson:"phone_number"`
	Email         string   `bson:"email"`
	Document      string   `bson:"document"`
	FullName      string   `bson:"full_name"`
	TypeOfAccount []string `bson:"type_of_account"`
	Roles         []string `bson:"roles"`

	Active  bool `bson:"active"`
	Blocked bool `bson:"blocked"`
	Removed bool `bson:"removed"`

	LoginInOnlyDevice bool   `bson:"login_in_only_device"`
	AccessToken       string `bson:"access_token"`

	PasswordHash     string             `bson:"password_hash"`
	PasswordStatus   string             `bson:"password_status"`
	PasswordValidity string             `bson:"password_validity"`
	ResetPasswordAt  primitive.DateTime `bson:"reset_password_at"`

	TokenSetPassword string             `bson:"token_set_password"`
	PasswordSetedAt  primitive.DateTime `bson:"password_seted_at"`

	TokenPasswordRecovery string             `bson:"token_password_recovery"`
	PasswordRecoveredAt   primitive.DateTime `bson:"password_recovered_at"`

	TokenResetPassword string             `bson:"token_reset_password"`
	ResetedPasswordAt  primitive.DateTime `bson:"reseted_password_at"`

	LastLoginAt primitive.DateTime `bson:"last_login_at"`

	ConnectedToDevice bool               `bson:"connected_to_device"`
	ActiveDeviceId    string             `bson:"ActiveDeviceId"`
	ActiveDeviceAt    primitive.DateTime `bson:"activeDeviceAt"`

	RegisterOrigin string `bson:"register_origin"`
	AvatarPath     string `bson:"avatar_path"`

	PhoneIsVerified     bool               `bson:"phone_is_verified"`
	ValidatedPhoneAt    primitive.DateTime `bson:"validated_phone_at"`
	TokenValidatedPhone string             `bson:"token_validated_phone"`

	EmailIsVerified     bool               `bson:"email_is_verified"`
	ValidatedEmailAt    primitive.DateTime `bson:"validated_email_at"`
	TokenValidatedEmail bool               `bson:"token_validated_email"`

	CreateAt primitive.DateTime `bson:"create_at"`
	UpdateAt primitive.DateTime `bson:"update_at"`
}
