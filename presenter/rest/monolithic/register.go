package rest

//
// presenter => rest => monolithic => register.go
//

import (
	"BackEnd_Api/logger"
	"context"

	//accountMicroservicePb "BackEnd_Api/microservices/account/presenter/grpc/v1/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterRESTModules(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {

	// if err := accountMicroservicePb.RegisterAccountHandler(ctx, mux, conn); err != nil {
	// 	logger.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	// 	return err
	// }

	logger.Log.Info("Não existem rotas Rest na Api")

	return nil
}
