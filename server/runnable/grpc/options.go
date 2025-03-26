package grpc

import (
	"crypto/tls"
	"fmt"
	"math"
	"time"

	pbErr "github.com/cpyun/gyopls-core/errors"
	"github.com/cpyun/gyopls-core/server/runnable/grpc/interceptors/logging"
	requesttag "github.com/cpyun/gyopls-core/server/runnable/grpc/interceptors/request_tag"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

const (
	infinity                           = time.Duration(math.MaxInt64)
	defaultMaxMsgSize                  = 4 << 20
	defaultMaxConcurrentStreams        = 100000
	defaultKeepAliveTime               = 30 * time.Second
	defaultConnectionIdleTime          = 10 * time.Second
	defaultMaxServerConnectionAgeGrace = 10 * time.Second
	defaultMiniKeepAliveTimeRate       = 2
)

type Optionfunc func(*grpcOptions)

type grpcOptions struct {
	id                       string
	domain                   string
	addr                     string
	tls                      *tls.Config
	keepAlive                time.Duration
	timeout                  time.Duration
	maxConnectionAge         time.Duration
	maxConnectionAgeGrace    time.Duration
	maxConcurrentStreams     int
	maxMsgSize               int
	unaryServerInterceptors  []grpc.UnaryServerInterceptor
	streamServerInterceptors []grpc.StreamServerInterceptor
	startedHook              func()
	endHook                  func()
}

func WithIDOption(s string) Optionfunc {
	return func(o *grpcOptions) {
		o.id = s
	}
}

func WithDomainOption(s string) Optionfunc {
	return func(o *grpcOptions) {
		o.domain = s
	}
}

func WithAddrOption(s string) Optionfunc {
	return func(o *grpcOptions) {
		o.addr = s
	}
}

func WithTlsOption(tls *tls.Config) Optionfunc {
	return func(o *grpcOptions) {
		o.tls = tls
	}
}

func WithKeepAliveOption(t time.Duration) Optionfunc {
	return func(o *grpcOptions) {
		o.keepAlive = t
	}
}

func WithTimeoutOption(t time.Duration) Optionfunc {
	return func(o *grpcOptions) {
		o.keepAlive = t
	}
}

func WithMaxConnectionAgeOption(t time.Duration) Optionfunc {
	return func(o *grpcOptions) {
		o.maxConnectionAge = t
	}
}

func WithMaxConnectionAgeGraceOption(t time.Duration) Optionfunc {
	return func(o *grpcOptions) {
		o.maxConnectionAgeGrace = t
	}
}

func WithMaxConcurrentStreamsOption(i int) Optionfunc {
	return func(o *grpcOptions) {
		o.maxConcurrentStreams = i
	}
}

func WithMaxMsgSizeOption(i int) Optionfunc {
	return func(o *grpcOptions) {
		o.maxMsgSize = i
	}
}

func WithUnaryServerInterceptorsOption(u ...grpc.UnaryServerInterceptor) Optionfunc {
	return func(o *grpcOptions) {
		if o.unaryServerInterceptors == nil {
			o.unaryServerInterceptors = make([]grpc.UnaryServerInterceptor, 0)
		}
		o.unaryServerInterceptors = append(o.unaryServerInterceptors, u...)
	}
}

func WithStreamServerInterceptorsOption(u ...grpc.StreamServerInterceptor) Optionfunc {
	return func(o *grpcOptions) {
		if o.streamServerInterceptors == nil {
			o.streamServerInterceptors = make([]grpc.StreamServerInterceptor, 0)
		}
		o.streamServerInterceptors = append(o.streamServerInterceptors, u...)
	}
}

func defaultOptions() *grpcOptions {
	return &grpcOptions{
		addr:                  ":0",
		keepAlive:             defaultKeepAliveTime,
		timeout:               defaultConnectionIdleTime,
		maxConnectionAge:      infinity,
		maxConnectionAgeGrace: defaultMaxServerConnectionAgeGrace,
		maxConcurrentStreams:  defaultMaxConcurrentStreams,
		maxMsgSize:            defaultMaxMsgSize,
		unaryServerInterceptors: []grpc.UnaryServerInterceptor{
			requesttag.UnaryServerInterceptor(),
			ctxtags.UnaryServerInterceptor(),
			opentracing.UnaryServerInterceptor(),
			logging.UnaryServerInterceptor(),
			prometheus.UnaryServerInterceptor,
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(customRecovery("", ""))),
		},
		streamServerInterceptors: []grpc.StreamServerInterceptor{
			requesttag.StreamServerInterceptor(),
			ctxtags.StreamServerInterceptor(),
			opentracing.StreamServerInterceptor(),
			logging.StreamServerInterceptor(),
			prometheus.StreamServerInterceptor,
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(customRecovery("", ""))),
		},
		startedHook: func() {
			fmt.Printf("[gRPC] Server listening on %s \r\n", ":0")
		},
		endHook: func() {
			fmt.Println("[gRPC] Server will be shutdown gracefully")
		},
	}
}

func customRecovery(id, domain string) recovery.RecoveryHandlerFunc {
	return func(p interface{}) (err error) {
		fmt.Printf("panic triggered: %v \r\n", p)
		return pbErr.New(id, domain, pbErr.InternalServerError)
	}
}
