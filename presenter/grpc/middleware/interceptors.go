package middleware

//
// presenter => grpc => middleware => interceptors.go
//

import (
	"BackEnd_Api/config"
	jwt "BackEnd_Api/helpers/jwt"
	account_microservice "BackEnd_Api/microservices/account"

	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddInterceptors(config *config.Config, logger *zap.Logger, tracer opentracing.Tracer, opts []grpc.ServerOption) []grpc.ServerOption {

	grpc_zap.ReplaceGrpcLoggerV2(logger)

	recoveryOptions := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(grpcPanicsRecovery),
	}

	dur, _ := time.ParseDuration(config.Auth.DurationInMinutes + "m")
	jwtManager := jwt.NewJwtManager(config.Auth.Secret, dur)

	controlledRoles := account_microservice.GetAccountControlled()

	authMiddleware := NewAuthInterceptor(jwtManager, controlledRoles)

	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_recovery.UnaryServerInterceptor(recoveryOptions...),

		grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),

		grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(logger),

		grpc_auth.UnaryServerInterceptor(authMiddleware.Unary),

		grpc_validator.UnaryServerInterceptor(),
	))

	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_recovery.StreamServerInterceptor(recoveryOptions...),
		grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),

		grpc_opentracing.StreamServerInterceptor(),

		grpc_prometheus.StreamServerInterceptor,

		grpc_auth.StreamServerInterceptor(authMiddleware.Stream),

		grpc_zap.StreamServerInterceptor(logger),

		grpc_validator.StreamServerInterceptor(),
	))

	return opts
}

func grpcPanicsRecovery(in interface{}) error {
	return status.Errorf(codes.Unknown, "Erro desconhecido")
}
