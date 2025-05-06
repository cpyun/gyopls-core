package http

import (
	"fmt"
	"net/http"
	"time"
)

// Option 参数设置类型
type OptionFunc func(*httpOptions)

type httpOptions struct {
	addr         string
	cert         *cert
	handler      http.Handler
	readTimeout  time.Duration
	writeTimeout time.Duration
	startedHook  func()
	endHook      func()
}

type cert struct {
	certFile, keyFile string
}

func setDefaultOption() httpOptions {
	h := httpOptions{
		addr:         ":8080",
		handler:      http.NotFoundHandler(),
		readTimeout:  60 * time.Second,
		writeTimeout: 60 * time.Second,
	}
	h.startedHook = func() {
		fmt.Println("[http] server startup successful")
	}

	return h
}

// WithAddr 设置addr
func WithAddr(s string) OptionFunc {
	return func(o *httpOptions) {
		o.addr = s
	}
}

// WithTlsOption 设置cert
func WithTlsOption(certFile, keyFile string) OptionFunc {
	return func(o *httpOptions) {
		o.cert = &cert{
			certFile: certFile,
			keyFile:  keyFile,
		}
	}
}

// WithHandler 设置handler
func WithHandler(handler http.Handler) OptionFunc {
	return func(o *httpOptions) {
		o.handler = handler
	}
}

// WithReadTimeout 设置读超时
func WithReadTimeout(d time.Duration) OptionFunc {
	return func(o *httpOptions) {
		o.readTimeout = d
	}
}

// WithWriteTimeout 设置写超时
func WithWriteTimeout(d time.Duration) OptionFunc {
	return func(o *httpOptions) {
		o.writeTimeout = d
	}
}

// WithStartedHook 设置启动回调函数
func WithStartedHook(f func()) OptionFunc {
	return func(o *httpOptions) {
		o.startedHook = f
	}
}

// WithEndHook 设置结束回调函数
func WithEndHook(f func()) OptionFunc {
	return func(o *httpOptions) {
		o.endHook = f
	}
}
