package server

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/cpyun/gyopls-core/logger"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	services               map[string]Runnable
	mutex                  sync.Mutex
	errChan                chan error
	waitForRunnable        sync.WaitGroup
	errGroup               errgroup.Group
	internalCtx            context.Context
	internalCancel         context.CancelFunc
	internalProceduresStop chan struct{}
	shutdownCtx            context.Context
	shutdownCancel         context.CancelFunc
	logger                 *logger.Logger
	opts                   options
}

// New 实例化
func New(opts ...OptionFunc) *Server {
	s := &Server{
		services:               make(map[string]Runnable),
		errChan:                make(chan error),
		internalProceduresStop: make(chan struct{}),
	}
	s.opts = setDefaultOptions()
	s.withOptions(opts...)
	return s
}

func (e *Server) withOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(&e.opts)
	}
}

// Add 添加 runnable
func (e *Server) Add(r ...Runnable) {
	if e.services == nil {
		e.services = make(map[string]Runnable)
	}
	for _, v := range r {
		e.services[v.String()] = v
	}
}

// Start 启动 runnable
func (e *Server) Start(ctx context.Context) (err error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.internalCtx, e.internalCancel = context.WithCancel(context.Background())
	defer func() {
		err = e.shutdownStopComplete(err)
	}()
	e.errChan = make(chan error)

	for _, v := range e.services {
		if !v.Attempt() {
			//先判断是否可以启动
			return errors.New("can't accept new runnable as stop procedure is already engaged")
		}
	}
	//按顺序启动
	for k := range e.services {
		e.startRunnable(e.services[k])
	}
	e.waitForRunnable.Wait()
	if err = e.errGroup.Wait(); err != nil {
		e.errChan <- err
	}
	select {
	case <-ctx.Done():
		return nil
	case err = <-e.errChan:
		return err
	}
}

func (e *Server) startRunnable(r Runnable) {
	e.waitForRunnable.Add(1)
	e.errGroup.Go(func() error {
		defer e.waitForRunnable.Done()
		if err := r.Start(e.internalCtx); err != nil {
			e.errChan <- err
			return err
		}
		return nil
	})
}

func (e *Server) shutdownStopComplete(err error) error {
	stopComplete := make(chan struct{})
	defer close(stopComplete)
	stopErr := e.engageStopProcedure(stopComplete)
	if stopErr != nil {
		if err != nil {
			err = fmt.Errorf("%s, %w", stopErr.Error(), err)
		} else {
			err = stopErr
		}
	}
	return err
}

func (e *Server) engageStopProcedure(stopComplete <-chan struct{}) error {
	if e.opts.gracefulShutdownTimeout > 0 {
		e.shutdownCtx, e.shutdownCancel = context.WithTimeout(context.Background(), e.opts.gracefulShutdownTimeout)
	} else {
		e.shutdownCtx, e.shutdownCancel = context.WithCancel(context.Background())
	}
	defer e.shutdownCancel()
	close(e.internalProceduresStop)
	e.internalCancel()

	go func() {
		for {
			select {
			case err, ok := <-e.errChan:
				if ok {
					e.logger.Error("error received after stop sequence was engaged", "error", err.Error())
				}
			case <-stopComplete:
				return
			}
		}
	}()

	return e.waitForRunnableToEnd()
}

func (e *Server) waitForRunnableToEnd() error {
	if e.opts.gracefulShutdownTimeout == 0 {
		go func() {
			e.waitForRunnable.Wait()
			e.shutdownCancel()
		}()
	}
	select {
	case <-e.shutdownCtx.Done():
		if err := e.shutdownCtx.Err(); err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			return fmt.Errorf(
				"failed waiting for all runnables to end within grace period of %s: %w",
				e.opts.gracefulShutdownTimeout, err)
		}
	}

	return nil
}

func (e *Server) Shutdown(ctx context.Context) (err error) {
	return e.shutdownStopComplete(err)
}
