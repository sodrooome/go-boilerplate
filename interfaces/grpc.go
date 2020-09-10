package interfaces

import (
	domain "backend-project/domain"
)

// function for initialize grpc authentication
func GrpcAuthentication(Key string) domain.Grpc {
	return domain.Grpc{Key: Key}
}

// function for handling grpc product's
func ProductGrpcHandler(ProductGrpc domain.ProductGrpcHandler) *domain.ProductGrpcHandler {
	return ProductGrpcHandler(ProductGrpc)
}