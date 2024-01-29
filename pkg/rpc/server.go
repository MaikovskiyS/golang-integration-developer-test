package rpc

import (
	"runtime/debug"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewServer(logger *zap.Logger) *grpc.Server {
	decider := func(_ string, err error) bool {
		return true
	}

	if logger.Level() > zapcore.WarnLevel {
		decider = func(_ string, err error) bool {
			return err != nil
		}
	}

	return grpc.NewServer(
		grpc.StreamInterceptor(middleware.ChainStreamServer(
			grpcValidator.StreamServerInterceptor(),
			grpcZap.StreamServerInterceptor(logger, grpcZap.WithDecider(decider)),
			grpcRecovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			grpcValidator.UnaryServerInterceptor(),
			grpcZap.UnaryServerInterceptor(logger, grpcZap.WithDecider(decider)),
			grpcRecovery.UnaryServerInterceptor(
				grpcRecovery.WithRecoveryHandler(
					func(p interface{}) error {
						return status.Errorf(codes.Internal, "panic: %v: stack: %s", p, string(debug.Stack()))
					},
				),
			),
		)),
	)
}
