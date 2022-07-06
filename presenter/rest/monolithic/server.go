package rest

//
// presenter => rest => monolithic => server.go
//

import (
	"BackEnd_Api/config"
	"BackEnd_Api/logger"
	"BackEnd_Api/presenter/rest/middleware"

	"context"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func RunRestServer(lc fx.Lifecycle, config *config.Config, conn *grpc.ClientConn) {
	Ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	if err := RegisterRESTModules(Ctx, mux, conn); err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr: config.Rest.Host + ":" + config.Rest.Port,

		Handler: http.TimeoutHandler(
			middleware.AddCors(middleware.AddRequestID(middleware.AddLogger(logger.Log, mux))),
			time.Second*time.Duration(config.Rest.RequestTimeout),
			"Prazo de contexto excedido",
		),

		WriteTimeout: time.Second * time.Duration(config.Rest.RequestTimeout),
	}

	logger.Log.Info("iniciando o gateway HTTP/REST...")

	lc.Append(fx.Hook{

		OnStart: func(ctx context.Context) error {

			logger.Log.Info("iniciando o gateway HTTP/REST...")
			go func() {
				_ = srv.ListenAndServe()
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Log.Info("Desligamento seguro do servidor REST")
			_ = srv.Shutdown(Ctx)
			return nil
		},
	})
}
