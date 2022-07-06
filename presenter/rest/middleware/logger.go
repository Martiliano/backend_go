package middleware

//
// presenter => rest => middleware => logger.go
//

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

func AddLogger(logger *zap.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Header.Get("X-Liveness-Probe") == "Healthz" {
			h.ServeHTTP(w, r)
			return
		}

		id := GetReqID(ctx)

		var scheme string
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
		proto := r.Proto
		method := r.Method
		remoteAddr := r.RemoteAddr
		userAgent := r.UserAgent()
		uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")

		var request interface{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			zap.Error(err)
		}

		if request == nil {
			request = "O objeto de solicitação é nulo"
		}

		logger.Debug("solicitação iniciada",
			zap.String("request-id", id),
			zap.String("http-scheme", scheme),
			zap.String("http-proto", proto),
			zap.String("http-method", method),
			zap.String("remote-addr", remoteAddr),
			zap.String("user-agent", userAgent),
			zap.String("uri", uri),
			zap.ByteString("request-body", []byte(fmt.Sprintf("%v", request.(interface{})))),
		)

		t1 := time.Now()
		h.ServeHTTP(w, r)

		logger.Debug("solicitação concluída",
			zap.String("request-id", id),
			zap.String("http-scheme", scheme),
			zap.String("http-proto", proto),
			zap.String("http-method", method),
			zap.String("remote-addr", remoteAddr),
			zap.String("user-agent", userAgent),
			zap.String("uri", uri),
			zap.Float64("elapsed-ms", float64(time.Since(t1).Nanoseconds())/1000000.0),
		)
	})
}
