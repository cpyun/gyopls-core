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
		fmt.Printf("[http] server listening on %s \r\n", h.addr)
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
func WithReadTimeout(d int) OptionFunc {
	return func(o *httpOptions) {
		o.readTimeout = time.Duration(d)
	}
}

// WithWriteTimeout 设置写超时
func WithWriteTimeout(d int) OptionFunc {
	return func(o *httpOptions) {
		o.writeTimeout = time.Duration(d)
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

// WithCert 设置cert
//
// Deprecated: Set tls cert.
// punctuation properly. Use WithTlsOption instead.
func WithCert(s string) OptionFunc {
	return func(o *httpOptions) {
		o.cert = &cert{
			certFile: s,
		}
	}
}

// WithKey 设置key
//
// Deprecated: Set tls key.
// punctuation properly. Use WithKey instead.
func WithKey(s string) OptionFunc {
	return func(o *httpOptions) {
		o.cert = &cert{
			keyFile: s,
		}
	}
}
