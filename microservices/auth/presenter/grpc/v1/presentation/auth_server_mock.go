package auth

//
// microservices => auth => presenter => grpc => v1 => presentation => auth_server_mock.go
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
	auth_repository_mongo "BackEnd_Api/microservices/auth/external/db/mongo/implementation"
	pb "BackEnd_Api/microservices/auth/presenter/grpc/v1/pb"
	auth_usecase "BackEnd_Api/microservices/auth/usecase"
)

func CreateNewAuthServer_Tester(done chan<- bool, quit <-chan bool) {

	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel, // TODO: Pegue este nível da configuração
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

	auth_repository_mongo := auth_repository_mongo.NewAuthMongoRepository(db, ctx, "BackEnd_Api_test")
	account_repository_mongo := account_repository_mongo.NewAccountMongoRepository(db, ctx, "BackEnd_Api_test")
	config := config.GetConfig()

	usecase := auth_usecase.NewAuthUseCase(auth_repository_mongo, account_repository_mongo, config)

	authMicroSevice := NewAuthServer(usecase)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7070))

	if err != nil {
		log.Printf("Erro ao abrir porta tcp : " + error.Error(err))
		return
	}

	grpc := grpc.NewServer()
	pb.RegisterAuthServer(grpc, authMicroSevice)

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
