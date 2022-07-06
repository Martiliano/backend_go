package middleware

//
// presenter => grpc => middleware => auth.go
//

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	jwt "BackEnd_Api/helpers/jwt"
	auth "BackEnd_Api/microservices/auth/rules"
)

type AuthInterceptor struct {
	jwtManager      *jwt.JwtManager
	controlledRoles map[string]auth.Controlled
}

func NewAuthInterceptor(jwtManager *jwt.JwtManager, controlledRoles map[string]auth.Controlled) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, controlledRoles}
}

func (interceptor *AuthInterceptor) Unary(ctx context.Context) (context.Context, error) {

	method, _ := grpc.Method(ctx)

	accessibleRoles, ok := interceptor.controlledRoles[method]
	if !ok {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadados não são fornecidos")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token não foi fornecido")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "o token de acesso é inválido: %v", err)
	}

	roles := ""
	for _, x := range claims.Roles {
		roles = roles + x + ","
	}

	for _, role := range accessibleRoles.Roles {

		for _, s := range claims.Roles {
			if s == role {
				ctx = metadata.AppendToOutgoingContext(ctx, "Id", claims.Id, "TypeOfAccount", strings.Join(claims.TypeOfAccount, ","), "Roles", strings.Join(claims.Roles, ","))
				return ctx, nil
			}
		}

	}

	return nil, status.Error(codes.PermissionDenied, "sem permissão para acessar este RPC")
}

func (interceptor *AuthInterceptor) Stream(ctx context.Context) (context.Context, error) {
	method, _ := grpc.Method(ctx)

	accessibleRoles, ok := interceptor.controlledRoles[method]
	if !ok {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadados não são fornecidos")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token não foi fornecido")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "o token de acesso é inválido: %v", err)
	}

	roles := ""
	for _, x := range claims.Roles {
		roles = roles + x + ","
	}

	for _, role := range accessibleRoles.Roles {

		for _, s := range claims.Roles {
			if s == role {
				ctx = metadata.AppendToOutgoingContext(ctx, "Id", claims.Id, "TypeOfAccount", strings.Join(claims.TypeOfAccount, ","), "Roles", strings.Join(claims.Roles, ","))
				return ctx, nil
			}
		}

	}

	return nil, status.Error(codes.PermissionDenied, "sem permissão para acessar este RPC")
}
