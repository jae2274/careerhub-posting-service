package utils

import (
	"github.com/jae2274/goutils/mw/grpcmw"
	"google.golang.org/grpc"
)

func Middlewares() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpcmw.GetTraceIdUnaryMW(),
			grpcmw.LogErrUnaryMW(),
		),
		grpc.ChainStreamInterceptor(
			grpcmw.GetTraceIdStreamMW(),
			grpcmw.LogErrStreamMW(),
		),
	}
}
