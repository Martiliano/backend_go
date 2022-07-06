package account

//
// microservices => account => microservice.go
//

import (
	config "BackEnd_Api/config"
	email_service "BackEnd_Api/helpers/email/gomailv2"
	account_repository_mongo "BackEnd_Api/microservices/account/external/db/mongo/implementation"
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	presentation "BackEnd_Api/microservices/account/presenter/grpc/v1/presentation"
	account_usecase "BackEnd_Api/microservices/account/usecase"
	"strconv"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewAccountMicroService(config *config.Config, db *mongo.Client) pb.AccountServer {
	ctx := context.Background()

	email := email_service.NewEmailService()
	email.Initialize()

	smtpSecure := false
	if config.Email.SmtpSecure == "ssl/tls" {
		smtpSecure = true
	}

	smtpPort, _ := strconv.Atoi(config.Email.SmtpPort)

	email.SetSmtpServer(config.Email.SmtpServer, smtpSecure, smtpPort)
	email.SetEmailCredentials(config.Email.User, config.Email.Password)
	email.SendEmailFrom(config.Email.Account, config.Email.AccountName)

	repository_mongo := account_repository_mongo.NewAccountMongoRepository(db, ctx, config.Mongo.DbName)

	usecase := account_usecase.NewAccountUseCase(repository_mongo, email, config, config.Auth.CreateSecret)

	return presentation.NewAccountServer(usecase)
}
