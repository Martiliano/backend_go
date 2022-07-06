package db

import (
	config "BackEnd_Api/config"
	"BackEnd_Api/logger"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func NewConnection(config *config.Config) *mongo.Client {
	ctx := context.Background()

	connectOptions := options.Client().ApplyURI(config.Mongo.Host + ":" + config.Mongo.Port) // mongodb://127.0.0.1:27017/
	db, err := mongo.Connect(ctx, connectOptions)

	if err != nil {
		logger.Log.Fatal("Tentativa de conexão com o MongoDB falhou", zap.Error(err))
		panic(err)
	}

	err = db.Ping(ctx, nil)

	if err != nil {
		logger.Log.Fatal("Teste de conexão com o MongoDB falhou ", zap.Error(err))
		panic(err)
	}

	logger.Log.Info("Conexão com o MongoDB estabelecida")

	return db
}
