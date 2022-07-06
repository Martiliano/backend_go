package grpcClient

//
// client => grpc => client.go
//

import (
	"BackEnd_Api/config"
	"BackEnd_Api/logger"
	"BackEnd_Api/metrics/opentracing/tracer/jaeger"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"context"
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientContext struct {
	Ctx context.Context
}

func GetGrpcClientConnection(config *config.Config) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(grpc_opentracing.WithTracer(jaeger.Tracer)),
		),
		grpc.WithStreamInterceptor(grpc_opentracing.StreamClientInterceptor(grpc_opentracing.WithTracer(jaeger.Tracer))),
	)

	// //append grpc insecure
	// opts = append(opts,
	// 	grpc.WithInsecure(),
	// )

	logger.Log.Info("Configurando os certificados do cliente...")

	pemClientCA, err := ioutil.ReadFile(config.ServerCertificate.CaCertFile)
	if err != nil {
		logger.Log.Fatal("não é possível ler o CA do servidor", zap.Error(err))
		panic(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		logger.Log.Fatal("falha ao adicionar o certificado da CA do cliente", zap.Error(err))
		panic(err)
	}

	clientCert, err := tls.LoadX509KeyPair(config.ServerCertificate.ServerCertFile,
		config.ServerCertificate.ServerKeyFile)
	if err != nil {
		logger.Log.Fatal("falha ao carregar o certificado e a chave privada do cliente", zap.Error(err))
		panic(err)
	}

	configCredentials := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		ClientCAs:    certPool,
	}

	tlsCredentials := credentials.NewTLS(configCredentials)

	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Grpc.RequestTimeout))
	defer cancel()

	conn, err := grpc.DialContext(ctx, config.Grpc.Host+":"+config.Grpc.Port, opts...)
	if err != nil {
		logger.Log.Fatal("o cliente não conectou", zap.Error(err))
	}
	return conn
}
