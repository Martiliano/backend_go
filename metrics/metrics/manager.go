package metrics

//
// metrics => metrics => manager.go
//

import (
	grpcErrors "BackEnd_Api/presenter/grpc/errors"
	"context"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

type Manager struct {
	metrics Metrics
}

func NewMetricsManager(metrics Metrics) *Manager {
	return &Manager{metrics: metrics}
}

func (im *Manager) Metrics(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	var status = http.StatusOK
	if err != nil {
		status = grpcErrors.MapGRPCErrCodeToHTTPStatus(grpcErrors.ParseGRPCErrStatusCode(err))
	}
	im.metrics.ObserveResponseTime(status, info.FullMethod, info.FullMethod, time.Since(start).Seconds())
	im.metrics.IncHits(status, info.FullMethod, info.FullMethod)

	return resp, err
}
