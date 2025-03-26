package http

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/cpyun/gyopls-core/server"
)

type httpApt struct {
	//ctx     context.Context
	name    string
	srv     *http.Server
	started bool
	opts    httpOptions
}

// Options 设置参数
func (e *httpApt) applyOptions(opts ...OptionFunc) {
	for _, o := range opts {
		o(&e.opts)
	}
}

func (e *httpApt) String() string {
	return e.name
}

// Start 开始
func (e *httpApt) Start(ctx context.Context) (err error) {
	e.srv = &http.Server{
		Addr:         e.opts.addr,
		Handler:      e.opts.handler,
		ReadTimeout:  e.opts.readTimeout,
		WriteTimeout: e.opts.writeTimeout,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}
	e.started = true
	if e.opts.startedHook != nil {
		e.opts.startedHook()
	}
	if e.opts.endHook != nil {
		e.srv.RegisterOnShutdown(e.opts.endHook)
	}

	// 启动
	if e.opts.cert != nil {
		err = e.srv.ListenAndServeTLS(e.opts.cert.certFile, e.opts.cert.keyFile)
	} else {
		err = e.srv.ListenAndServe()
	}
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("%s Server start error: %s \r\n", e.name, err.Error())
	}

	return nil
}

// Attempt 判断是否可以启动
func (e *httpApt) Attempt() bool {
	return !e.started
}

// Shutdown 停止
func (e *httpApt) Shutdown(ctx context.Context) error {
	<-ctx.Done()
	return e.srv.Shutdown(ctx)
}

// New 实例化
func New(name string, opts ...OptionFunc) server.Runnable {
	s := &httpApt{
		name: name,
		opts: setDefaultOption(),
	}

	s.applyOptions(opts...)
	return s
}
