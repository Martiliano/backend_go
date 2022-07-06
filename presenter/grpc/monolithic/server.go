package grpc

//
// presenter => grpc => monolithic => server.go
//

import (
	"BackEnd_Api/config"
	"BackEnd_Api/logger"
	"BackEnd_Api/presenter/grpc/middleware"

	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func InitGrpcBeforeServing(config *config.Config, tracer opentracing.Tracer) (*grpc.Server, net.Listener) {

	logger.Log.Info("gRPC Server Host: " + config.Grpc.Host)
	logger.Log.Info("gRPC Server Port: " + config.Grpc.Port)

	logger.Log.Info("Client CA Cert File: " + config.ServerCertificate.CaCertFile)
	logger.Log.Info("Server Cert File: " + config.ServerCertificate.ServerCertFile)
	logger.Log.Info("Server Key File: " + config.ServerCertificate.ServerKeyFile)

	listen, err := net.Listen("tcp", config.Grpc.Host+":"+config.Grpc.Port)
	if err != nil {
		logger.Log.Fatal("não é possível inicializar o servidor grpc", zap.Error(err))
		panic(err)
	}

	var opts []grpc.ServerOption

	opts = middleware.AddInterceptors(config, logger.Log, tracer, opts)

	logger.Log.Info("Configurando os certificados do servidor...")

	pemClientCA, err := ioutil.ReadFile(config.ServerCertificate.CaCertFile)
	if err != nil {
		logger.Log.Fatal("não é possível ler o CA dos clientes", zap.Error(err))
		panic(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		logger.Log.Fatal("falha ao adicionar o certificado da CA do cliente", zap.Error(err))
		panic(err)
	}

	serverCert, err := tls.LoadX509KeyPair(config.ServerCertificate.ServerCertFile,
		config.ServerCertificate.ServerKeyFile)
	if err != nil {
		logger.Log.Fatal("falha ao carregar o certificado e a chave privada do servidor", zap.Error(err))
		panic(err)
	}

	configCredentials := &tls.Config{
		Certificates:       []tls.Certificate{serverCert},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          certPool,
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	tlsCredentials := credentials.NewTLS(configCredentials)

	opts = append(opts, grpc.Creds(tlsCredentials))

	server := grpc.NewServer(opts...)
	return server, listen
}

func RunGRPCServer(lc fx.Lifecycle, server *grpc.Server, listener net.Listener) error {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				reflection.Register(server)
				grpc_prometheus.EnableHandlingTimeHistogram()
				grpc_prometheus.Register(server)
				http.Handle("/metrics", promhttp.Handler())

				logger.Log.Info("Prometheus endpoint de métricas registrado em /metrics")

				logger.Log.Info("Iniciando o servidor HTTP2/gRPC...")

				go server.Serve(listener)
				return nil
			},
			OnStop: func(ctx context.Context) error {

				logger.Log.Info("Parando o servidor gRPC...")
				server.GracefulStop()
				return nil
			},
		},
	)

	return nil
}
