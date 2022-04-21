package logger

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func UnaryServerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	grpc_zap.ReplaceGrpcLoggerV2(logger)

	opts := grpc_zap.WithLevels(
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel
			case codes.Internal:
				l = zapcore.ErrorLevel
			default:
				l = zapcore.DebugLevel
			}
			return l
		},
	)

	return grpc_zap.UnaryServerInterceptor(logger, opts)
}
