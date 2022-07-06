package healthcheck

//
// microservices => metrics => metrics => healtcheck =>  health-check.go
//

import (
	"BackEnd_Api/config"
	mongohealth "BackEnd_Api/metrics/healthcheck/db/mongo"

	"net/http"

	"github.com/dimiro1/health"
	"go.mongodb.org/mongo-driver/mongo"
)

func RunHealthCheckServer(config *config.Config, db *mongo.Client) error {

	mongodb := mongohealth.NewMongoChecker(db, config.Mongo.DbName)
	handler := health.NewHandler()
	mongodb.Check()
	handler.AddChecker("Mongo", mongodb)
	http.Handle("/health/", handler)
	return http.ListenAndServe(config.HealthCheck.Host+":"+config.HealthCheck.Port, nil)
}

func HealthCheckRunner(config *config.Config, db *mongo.Client) error {
	go func() {
		_ = RunHealthCheckServer(config, db)
	}()
	return nil
}
