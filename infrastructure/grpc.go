package infrastructure

import (
	"backend-project/domain"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var g *domain.Grpc

func Authenticate(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	key, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "Missing something like key or metadata")
	}
	if len(key["authorization"]) != 1 {
		return nil, grpc.Errorf(codes.Unauthenticated, "Invalid key of authorization")
	}
	auth := key["authorization"][0]
	if auth != g.Key {
		return nil, grpc.Errorf(codes.Unauthenticated, "Invalid of meta authorization")
	}
	return handler(ctx, request)
}
