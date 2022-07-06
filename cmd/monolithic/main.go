package main

// fotos/amadora-1979

//
// cmd => monolithic => main.go
//

//
// COMANDO PARA FAZER O BUILD DA APLICACAO:
//
// go build -o bin/server -buildmode pie ./cmd/monolithic
//
// COMANDO PARA DAR PERMISSAO DE EXECUÇÃO:
//
// chmod a+x bin/server
//
// COMANDO PARA EXECUTAR O IURIS
//
// ./bin/server
//

// ENDPOINTS

// Jaeger UI:

// https://localhost:16686

// Health Trace:

// http://localhost:8083/health/

// Prometheus UI:

// http://localhost:9090

// Prometheus UI Metrics:

// http://localhost:9090/metrics

// Grpc Server:

// http://localhost:8080

// Graphql Server:

// http://localhost:8081

// Rest Server:

// http://localhost:8082

import (
	"go.uber.org/zap"

	logger "BackEnd_Api/logger"
)

func main() {
	logger.Init(logger.Config{
		LogLevel:    zap.DebugLevel,
		Development: false,
	})
	GetApp().Run()
}
