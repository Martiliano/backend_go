package middleware

//
// presenter => graphql => middleware => ctx-change.go
//

import (
	"net/http"

	"google.golang.org/grpc/metadata"
)

func ChangeContext(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		request = request.WithContext(metadata.AppendToOutgoingContext(request.Context(), "Authorization", request.Header.Get("Authorization")))
		h.ServeHTTP(writer, request)
	})
}
