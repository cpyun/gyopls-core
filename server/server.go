package server

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	servers     map[string]Runnable
	lock        sync.RWMutex
	eg          *errgroup.Group
	internalCtx context.Context
	opts        options
}

// New 实例化
func New(opts ...OptionFunc) *Server {
	s := &Server{
		servers: make(map[string]Runnable),
		opts:    setDefaultOptions(),
	}
	s.applyOptions(opts...)
	return s
}

func (e *Server) applyOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(&e.opts)
	}
}

// Add 添加 runnable
func (e *Server) Add(r ...Runnable) {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.servers == nil {
		e.servers = make(map[string]Runnable)
	}

	for _, v := range r {
		e.servers[v.String()] = v
	}
}

// Start 启动 runnable
func (e *Server) Start(ctx context.Context) (err error) {
	e.eg, e.internalCtx = errgroup.WithContext(ctx)
	// 启动
	e.lock.RLock()
	for _, srv := range e.servers {
		e.eg.Go(func() error {
			return e.startRunnable(srv)
		})
	}
	e.lock.RUnlock()

	//
	if err = e.eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (e *Server) startRunnable(r Runnable) error {
	//判断是否可以启动
	if !r.Attempt() {
		return fmt.Errorf("[%s] can't accept new runnable as stop procedure is already engaged", r.String())
	}

	if err := r.Start(e.internalCtx); err != nil {
		return err
	}
	return nil
}

func (e *Server) shutdownStopComplete(ctx context.Context) error {
	fmt.Printf("waiting for all runnables to end within grace period of %.2f second\r\n", e.opts.gracefulShutdownTimeout.Seconds())
	shutdownCtx, cancel := context.WithTimeout(ctx, e.opts.gracefulShutdownTimeout)
	defer cancel()
	return e.engageStopProcedure(shutdownCtx)
}

// 启动停止程序
func (e *Server) engageStopProcedure(ctx context.Context) error {
	var wg sync.WaitGroup
	for _, srv := range e.servers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := srv.Shutdown(ctx); err != nil {
				fmt.Printf("[%s] server shutdown error: %s \r\n", srv.String(), err.Error())
			}
		}()
	}
	wg.Wait()

	return nil
}

func (e *Server) Shutdown(ctx context.Context) (err error) {
	return e.shutdownStopComplete(ctx)
}
