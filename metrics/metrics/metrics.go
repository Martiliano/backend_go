package metrics

//
// metrics => metrics => metrics.go
//

import (
	"BackEnd_Api/config"
	"BackEnd_Api/logger"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

type Metrics interface {
	IncHits(status int, method, path string)
	ObserveResponseTime(status int, method, path string, observeTime float64)
}

type PrometheusMetrics struct {
	HitsTotal prometheus.Counter
	Hits      *prometheus.CounterVec
	Times     *prometheus.HistogramVec
}

func CreateMetrics(config *config.Config) Metrics {
	var metrics PrometheusMetrics
	metrics.HitsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: config.Metrics.ServiceName + "_hits_total",
	})
	if err := prometheus.Register(metrics.HitsTotal); err != nil {
		logger.Log.Fatal("erro ao registrar hits do prometheus", zap.Error(err))
		return nil
	}
	metrics.Hits = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: config.Metrics.ServiceName + "_hits",
		},
		[]string{"status", "method", "path"},
	)
	if err := prometheus.Register(metrics.Hits); err != nil {
		logger.Log.Fatal("erro ao registrar hits do prometheus", zap.Error(err))
		return nil
	}
	metrics.Times = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: config.Metrics.ServiceName + "_times",
		},
		[]string{"status", "method", "path"},
	)
	if err := prometheus.Register(metrics.Times); err != nil {
		logger.Log.Fatal("erro ao registrar o prometheus", zap.Error(err))
		return nil
	}
	if err := prometheus.Register(prometheus.NewBuildInfoCollector()); err != nil {
		logger.Log.Fatal("erro ao registrar o prometheus", zap.Error(err))
		return nil
	}
	go func() {
		router := echo.New()
		router.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		if err := router.Start(config.Metrics.URL); err != nil {
			logger.Log.Fatal("não foi possível criar métricas", zap.Error(err))
		}
	}()
	logger.Log.Info("Métricas disponíveis na URL: ", zap.String("url", config.Metrics.URL),
		zap.String("Nome do Serviço:", config.Metrics.ServiceName))
	return &metrics
}

func (metrics *PrometheusMetrics) IncHits(status int, method, path string) {
	metrics.HitsTotal.Inc()
	metrics.Hits.WithLabelValues(strconv.Itoa(status), method, path).Inc()
}

func (metrics *PrometheusMetrics) ObserveResponseTime(status int, method, path string, observeTime float64) {
	metrics.Times.WithLabelValues(strconv.Itoa(status), method, path).Observe(observeTime)
}
