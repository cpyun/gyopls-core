package zap

import (
	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/logger/level"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	// handler *zap.Logger        // zap句柄
	handler *zap.SugaredLogger // sugared
	opts    zapOption
}

func (l *zapLog) init() {
	var allCore []zapcore.Core
	allCore = append(allCore, zapcore.NewCore(l.getZapEncoder(), l.getZapWriteSyncer(), l.opts.level))

	// Adding namespace
	fields := []zap.Field{}
	if l.opts.namespaceKey != "" {
		fields = append(fields, zap.Namespace(l.opts.namespaceKey))
	}

	// log, _ := zapConfig.Build(zap.AddCallerSkip(l.opts.callerSkipKey), zap.Fields(fields...))
	zapCore := zapcore.NewTee(allCore...)
	log := zap.New(zapCore,
		zap.AddCaller(),
		zap.AddCallerSkip(l.opts.callerSkipKey),
		zap.AddStacktrace(zap.ErrorLevel),
		zap.Fields(fields...),
	)

	// defer log.Sync() ??
	// l.handler = log
	l.handler = log.Sugar()
}

func (l *zapLog) getZapWriteSyncer() zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer(l.opts.output...)
}

func (l *zapLog) getZapEncoder() zapcore.Encoder {
	cfg := l.getZapConfig()
	if cfg.Encoding == "console" {
		return zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	}
	return zapcore.NewJSONEncoder(cfg.EncoderConfig)
}

func (l *zapLog) getZapConfig() zap.Config {
	var zapConfig zap.Config
	//mode
	switch l.opts.mode {
	case "debug":
		zapConfig = zap.NewDevelopmentConfig()
	default:
		zapConfig = zap.NewProductionConfig()
	}

	// Level
	zapConfig.Level.SetLevel(l.opts.level)

	// EncoderConfig
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if l.opts.timeFormat != "" {
		zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(l.opts.timeFormat)
	}
	return zapConfig
}

func (l *zapLog) With(fields ...any) contract.LoggerHandler {
	newHandler := l.handler.With(fields...)
	//
	return &zapLog{
		handler: newHandler,
		opts:    l.opts,
	}
}

func (l *zapLog) Log(level level.Level, msg string, args ...any) {
	lvl := loggerLevelToZapLevel(level)
	// data := l.parseZapFields(args...)

	l.handler.Logw(lvl, msg, args...)
	defer l.handler.Sync()
}

func (l *zapLog) String() string {
	return "zap"
}

func NewZap(opts ...OptionFunc) contract.LoggerHandler {
	l := &zapLog{
		opts: setDefaultOptions(),
	}
	for _, o := range opts {
		o(&l.opts)
	}

	l.init()
	return l
}

func zaplevelToLoggerLevel(l zapcore.Level) level.Level {
	switch l {
	case zapcore.DebugLevel:
		return level.DebugLevel
	case zapcore.InfoLevel:
		return level.InfoLevel
	case zapcore.WarnLevel:
		return level.WarnLevel
	case zapcore.ErrorLevel:
		return level.ErrorLevel
	case zapcore.DPanicLevel, zapcore.PanicLevel:
		return level.ErrorLevel
	case zapcore.FatalLevel:
		return level.FatalLevel
	default:
		return level.InfoLevel
	}
}

func loggerLevelToZapLevel(l level.Level) zapcore.Level {
	switch l {
	case level.TraceLevel, level.DebugLevel:
		return zapcore.DebugLevel
	case level.InfoLevel:
		return zapcore.InfoLevel
	case level.WarnLevel:
		return zapcore.WarnLevel
	case level.ErrorLevel:
		return zapcore.ErrorLevel
	case level.FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
