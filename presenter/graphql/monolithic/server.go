package graphql

//
// presenter => graphql => middleware => server.go
//

import (
	"context"
	"net/http"
	"time"

	"BackEnd_Api/config"
	"BackEnd_Api/logger"
	"BackEnd_Api/presenter/graphql/middleware"

	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunGraphqlServer(lc fx.Lifecycle, config *config.Config, conn *grpc.ClientConn) {
	Ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	if err := RegisterGraphqlModules(mux, conn); err != nil {
		logger.Log.Fatal("não é possível registrar módulos graphql", zap.Error(err))
	}
	http.Handle("/graphql", mux)
	srv := &http.Server{
		Addr:         config.Graphql.Host + ":" + config.Graphql.Port,
		WriteTimeout: time.Second * time.Duration(config.Graphql.RequestTimeout),

		Handler: http.TimeoutHandler(middleware.ChangeContext(middleware.AddCors(middleware.AddRequestID(
			middleware.AddLogger(logger.Log, mux)))), time.Second*time.Duration(config.Graphql.RequestTimeout), "Prazo de contexto excedido"),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			logger.Log.Info("iniciando o gateway HTTP/GRAPHQL...")
			go func() {
				_ = srv.ListenAndServe()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Log.Info("Desligamento seguro do servidor Graphql")
			_ = srv.Shutdown(Ctx)
			return nil
		},
	})
}
