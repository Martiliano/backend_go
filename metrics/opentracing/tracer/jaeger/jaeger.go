package jaeger

//
// metrics => opentracing => trace => jaeger => jaeger.go
//

import (
	"BackEnd_Api/config"
	"BackEnd_Api/logger"
	"io"
	"strconv"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
)

type Jaeger struct {
	Host        string
	ServiceName string
	LogSpans    bool
}

var (
	Tracer opentracing.Tracer
)

func InitJaeger(config *config.Config) (io.Closer, opentracing.Tracer) {

	logSpan, _ := strconv.ParseBool(config.Jaeger.LogSpans)

	jaegerCfgInstance := jaegerconfig.Configuration{
		ServiceName: config.Jaeger.ServiceName,
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans:           logSpan,
			LocalAgentHostPort: config.Jaeger.Host + ":" + config.Jaeger.Port,
		},
	}
	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegerconfig.Logger(jaegerlog.StdLogger),
		jaegerconfig.Metrics(metrics.NullFactory),
	)
	if err != nil {
		logger.Log.Fatal("não pode criar rastreador", zap.Error(err))
		panic(err)
	}

	logger.Log.Info("Jaeger rastreador iniciado")
	opentracing.SetGlobalTracer(tracer)
	return closer, tracer
}
