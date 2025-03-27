package server

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	services    map[string]Runnable
	mutex       sync.RWMutex
	errGroup    *errgroup.Group
	internalCtx context.Context
	opts        options
}

// New 实例化
func New(opts ...OptionFunc) *Server {
	s := &Server{
		services: make(map[string]Runnable),
		opts:     setDefaultOptions(),
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
	if e.services == nil {
		e.services = make(map[string]Runnable)
	}

	e.mutex.Lock()
	defer e.mutex.Unlock()

	for _, v := range r {
		e.services[v.String()] = v
	}
}

// Start 启动 runnable
func (e *Server) Start(ctx context.Context) (err error) {
	e.errGroup, e.internalCtx = errgroup.WithContext(ctx)
	// 启动
	e.mutex.RLock()
	for _, srv := range e.services {
		e.errGroup.Go(func() error {
			return e.startRunnable(srv)
		})
	}
	e.mutex.RUnlock()

	// 监听e.internalCtx.Done()，关闭所有Server
	e.errGroup.Go(func() error {
		return e.shutdownStopComplete()
	})
	//
	if err = e.errGroup.Wait(); err != nil {
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

func (e *Server) shutdownStopComplete() error {
	<-e.internalCtx.Done() // 等待 context 被取消

	fmt.Printf("waiting for all runnables to end within grace period of %.2f second\r\n", e.opts.gracefulShutdownTimeout.Seconds())
	shutdownCtx, cancel := context.WithTimeout(context.Background(), e.opts.gracefulShutdownTimeout)
	defer cancel()
	return e.engageStopProcedure(shutdownCtx)
}

// 启动停止程序
func (e *Server) engageStopProcedure(ctx context.Context) error {
	var wg sync.WaitGroup
	for _, srv := range e.services {
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
	e.errGroup.Go(func() error {
		return fmt.Errorf("server shutdown")
	})
	return nil
}
