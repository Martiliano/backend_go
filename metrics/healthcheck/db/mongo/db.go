package db

//
// microservices => metrics => metrics => healtcheck => db => mongo => db.go
//

import (
	"context"

	"github.com/dimiro1/health"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Checker struct {
	db     *mongo.Client
	dbName string
}

func NewMongoChecker(db *mongo.Client, dbName string) Checker {
	return Checker{
		db:     db,
		dbName: dbName,
	}
}

func (c Checker) Check() health.Health {

	h := health.NewHealth()

	if c.db == nil {
		h.Down().AddInfo("error", "Sem banco de dados")
		return h
	}

	err := c.db.Ping(context.TODO(), nil)

	if err != nil {
		h.Down().AddInfo("error", err.Error())
		return h
	}

	var commandResult bson.M
	command := bson.D{{"serverStatus", 1}}
	err = c.db.Database("test").RunCommand(context.TODO(), command).Decode(&commandResult)

	if err != nil {
		h.Down().AddInfo("error", err.Error())
		return h
	}

	h.Up().AddInfo("version", commandResult["version"])

	return h
}
