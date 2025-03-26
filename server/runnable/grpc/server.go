package grpc

import (
	"context"
	"fmt"
	"net"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	name    string
	srv     *grpc.Server
	started bool
	options grpcOptions
}

// New new grpc server
func New(name string, opts ...Optionfunc) *Server {
	s := &Server{
		name:    name,
		options: *defaultOptions(),
	}
	s.applyOptions(opts...)

	s.init()
	return s
}

// String string
func (e *Server) String() string {
	return e.name
}

func (e *Server) applyOptions(opts ...Optionfunc) {
	for _, o := range opts {
		o(&e.options)
	}
}

func (e *Server) init() {
	grpc.EnableTracing = false
	e.srv = grpc.NewServer(e.initGrpcServerOptions()...)
}

func (e *Server) Register(do func(server *Server)) {
	do(e)
	prometheus.Register(e.srv)
}

func (e *Server) initGrpcServerOptions() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(e.options.unaryServerInterceptors...)),
		grpc.StreamInterceptor(middleware.ChainStreamServer(e.options.streamServerInterceptors...)),
		grpc.MaxConcurrentStreams(uint32(e.options.maxConcurrentStreams)),
		grpc.MaxRecvMsgSize(e.options.maxMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime: e.options.keepAlive / defaultMiniKeepAliveTimeRate,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:                  e.options.keepAlive,
			Timeout:               e.options.timeout,
			MaxConnectionAge:      e.options.maxConnectionAge,
			MaxConnectionAgeGrace: e.options.maxConnectionAgeGrace,
		}),
	}
}

func (e *Server) Start(ctx context.Context) error {
	if !e.Attempt() {
		return fmt.Errorf("gRPC Server was started more than once. " +
			"This is likely to be caused by being added to a manager multiple times")
	}
	e.started = true
	e.options.startedHook()

	ts, err := net.Listen("tcp", e.options.addr)
	if err != nil {
		return fmt.Errorf("gRPC Server listening on %s failed: %w", e.options.addr, err)
	}

	// 启动服务
	if err = e.srv.Serve(ts); err != nil {
		return fmt.Errorf("gRPC Server start error: %s \r\n", err.Error())
	}

	return nil
}

func (e *Server) Attempt() bool {
	return !e.started
}

func (e *Server) Shutdown(ctx context.Context) error {
	<-ctx.Done()
	e.options.endHook()

	e.srv.GracefulStop()
	return nil
}
