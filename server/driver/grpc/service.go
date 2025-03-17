package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/cpyun/gyopls-core/server/driver/grpc/interceptors/logging"
	reqtags "github.com/cpyun/gyopls-core/server/driver/grpc/interceptors/request_tag"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
)

type Service struct {
	Connection  *grpc.ClientConn
	CallTimeout time.Duration
}

func (e *Service) Dial(
	endpoint string,
	callTimeout time.Duration,
	unary ...grpc.UnaryClientInterceptor) (err error) {
	fmt.Printf("configure service with endpoint: %s \r\n", endpoint)

	ctx, cancel := context.WithTimeout(context.Background(), callTimeout)
	defer cancel()

	if len(unary) == 0 {
		unary = defaultUnaryClientInterceptors()
	}
	e.Connection, err = grpc.DialContext(ctx,
		endpoint,
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(middleware.ChainStreamClient(defaultStreamClientInterceptors()...)),
		grpc.WithUnaryInterceptor(middleware.ChainUnaryClient(unary...)),
		grpc.WithDefaultCallOptions(grpc.WaitForReady(true), grpc.MaxCallRecvMsgSize(defaultMaxMsgSize)),
	)

	if err != nil {
		return fmt.Errorf("%w, connect gRPC service %s failed", err, endpoint)
	}
	return nil
}

func defaultUnaryClientInterceptors() []grpc.UnaryClientInterceptor {
	return []grpc.UnaryClientInterceptor{
		opentracing.UnaryClientInterceptor(),
		logging.UnaryClientInterceptor(),
		reqtags.UnaryClientInterceptor(),
	}
}

func defaultStreamClientInterceptors() []grpc.StreamClientInterceptor {
	return []grpc.StreamClientInterceptor{
		opentracing.StreamClientInterceptor(),
		logging.StreamClientInterceptor(),
		reqtags.StreamClientInterceptor(),
	}
}
