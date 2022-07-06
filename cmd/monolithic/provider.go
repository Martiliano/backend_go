package main

//
// cmd => monolithic => provider.go
//

import (
	grpcClient "BackEnd_Api/client/grpc"
	"BackEnd_Api/config"
	db "BackEnd_Api/external/db/mongo"
	"BackEnd_Api/metrics/opentracing/tracer/jaeger"
	"BackEnd_Api/presenter/cleanup"
	grpc "BackEnd_Api/presenter/grpc/monolithic"

	"go.uber.org/fx"
)

func GetProviderOptions() []fx.Option {
	return []fx.Option{
		config.ProviderFx,
		grpc.InitGrpcBeforeServingFx,
		db.DatabaseConnectionFx,
		cleanup.CleanupFx,
		jaeger.JaegerTracerFx,
		grpcClient.ConnectionFx,
	}
}
