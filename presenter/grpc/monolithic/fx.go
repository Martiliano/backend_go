package grpc

//
// presenter => grpc => monolithic => server.go
//

import "go.uber.org/fx"

var InitGrpcBeforeServingFx = fx.Options(
	fx.Provide(
		InitGrpcBeforeServing,
	),
)
