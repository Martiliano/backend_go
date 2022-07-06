package grpcClient

//
// client => grpc => fx.go
//

import "go.uber.org/fx"

var ConnectionFx = fx.Options(
	fx.Provide(
		GetGrpcClientConnection,
	),
)
