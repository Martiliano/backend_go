package graphql

//
// presenter => graphql => middleware => register.go
//

import (
	"BackEnd_Api/logger"

	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
	"google.golang.org/grpc"
)

func RegisterGraphqlModules(mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	logger.Log.Info("Não existem rotas GraphQl na Api")

	return nil
}
