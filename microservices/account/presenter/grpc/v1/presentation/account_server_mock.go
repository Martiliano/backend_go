package account

//
// microservices => account => presenter => grpc => v1 => presentation => account_server_mock.go
//
//

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"BackEnd_Api/config"
	"BackEnd_Api/logger"
	account_repository_mongo "BackEnd_Api/microservices/account/external/db/mongo/implementation"
	pb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"
	account_usecase "BackEnd_Api/microservices/account/usecase"
)

func CreateNewAccountServer_Tester(done chan<- bool, quit <-chan bool, messageBody *string) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	connectOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	db, err := mongo.Connect(ctx, connectOptions)
	defer db.Disconnect(ctx)

	if err != nil {
		log.Printf("Erro ao conectar ao MongoDb : " + error.Error(err))
		return
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		log.Printf("Erro ao testar a conexão ao MongoDb : " + error.Error(err))
		return
	}

	email := NewEmailMockService(messageBody)

	repository_mongo := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")
	config := config.GetConfig()

	usecase := account_usecase.NewAccountUseCase(repository_mongo, email, config, "secret_create_phrase")

	accountMicroSevice := NewAccountServer(usecase)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7070))

	if err != nil {
		log.Printf("Erro ao abrir porta tcp : " + error.Error(err))
		return
	}

	grpc := grpc.NewServer()
	pb.RegisterAccountServer(grpc, accountMicroSevice)

	wg := sync.WaitGroup{}
	wg.Add(1)

	done <- true

	go func() {
		<-quit

		log.Printf("Tentando executar o desligamento normal do servidor")
		cancel()

		grpc.GracefulStop()

		wg.Done()
	}()

	log.Println("iniciando o servidor grpc")

	err = grpc.Serve(lis)
	if err != nil {
		log.Fatalf("não poderia servir: %v", err)
	}

	wg.Wait()
	log.Println("desligamento limpo")

	done <- true

}
