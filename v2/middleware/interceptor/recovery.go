package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Recovery Recovery
func Recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("grpc err: %v", r)
				err = status.Errorf(codes.Unknown, "grpc err: %v", r)
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
