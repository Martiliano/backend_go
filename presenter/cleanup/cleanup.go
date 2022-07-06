package cleanup

//
// presenter => cleanup => cleanup.go
//

import (
	"BackEnd_Api/logger"
	"context"
	"io"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Config struct {
	GrpcServerConnection *grpc.Server
	GrpcClientConnection *grpc.ClientConn
	JaegerCloser         io.Closer
}

func GetCleanupConfig(
	GrpcServerConnection *grpc.Server,
	GrpcClientConnection *grpc.ClientConn,
	JaegerCloser io.Closer) *Config {
	return &Config{
		GrpcServerConnection: GrpcServerConnection,
		JaegerCloser:         JaegerCloser,
	}
}

func Cleanup(lc fx.Lifecycle, config *Config) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			logger.Log.Info(".......Iniciando o código de limpeza ......")

			config.GrpcServerConnection.GracefulStop()
			logger.Log.Info("conexão do servidor grpc fechada com sucesso")

			defer config.JaegerCloser.Close()
			logger.Log.Info("código de limpeza executado com sucesso")

			return nil
		},
	})

}
