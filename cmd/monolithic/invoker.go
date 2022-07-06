package main

//
// cmd => monolithic => invoker.go
//

import (
	"BackEnd_Api/metrics/healthcheck"
	"BackEnd_Api/presenter/cleanup"

	grpc "BackEnd_Api/presenter/grpc/monolithic"

	"go.uber.org/fx"
)

func GetInvokersOptions() fx.Option {
	return fx.Invoke(
		grpc.RunGRPCServer,
		grpc.RegisterGrpcModules,
		healthcheck.HealthCheckRunner,
		cleanup.Cleanup,
	)
}
