package zap

import (
	"fmt"

	"github.com/cpyun/gyopls-core/contract"
	"github.com/cpyun/gyopls-core/logger/level"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	handler *zap.Logger // zap句柄
	opts    zapOption
}

func (l *zapLog) init() {
	var allCore []zapcore.Core

	zapConfig := l.getZapConfig()
	encoder := l.getZapEncoder(zapConfig)
	writeSyncer := l.getZapWriteSyncer(zapConfig)
	allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapConfig.Level))
	//
	// logCore := zapcore.NewCore(
	// 	zapcore.NewConsoleEncoder(zapConfig.EncoderConfig),
	// 	zapcore.NewMultiWriteSyncer(zapcore.AddSync(writer)),
	// 	zapConfig.Level)

	fields := []zap.Field{}
	// Adding namespace
	if l.opts.namespaceKey != "" {
		fields = append(fields, zap.Namespace(l.opts.namespaceKey))
	}

	// log, _ := zapConfig.Build(zap.AddCallerSkip(l.opts.callerSkipKey), zap.Fields(fields...))
	zapCore := zapcore.NewTee(allCore...)
	log := zap.New(zapCore,
		zap.AddCaller(),
		zap.AddCallerSkip(l.opts.callerSkipKey),
		zap.AddStacktrace(zap.DPanicLevel),
		zap.Fields(fields...),
	)

	// defer log.Sync() ??
	l.handler = log
}

func (l *zapLog) getZapWriteSyncer(cfg zap.Config) zapcore.WriteSyncer {
	// return zapcore.AddSync(l.opts.out)
	var ws []zapcore.WriteSyncer
	for _, out := range cfg.OutputPaths {
		// cutter 需要注意 stderr|stdout 两种情况
		ws = append(ws, zapcore.AddSync(l.opts.cutter.SetFile(out)))
	}

	return zapcore.NewMultiWriteSyncer(ws...)
}

func (l *zapLog) getZapEncoder(cfg zap.Config) zapcore.Encoder {
	if cfg.Encoding == "console" {
		return zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	}
	return zapcore.NewJSONEncoder(cfg.EncoderConfig)
}

func (l *zapLog) getZapConfig() zap.Config {
	var zapConfig zap.Config
	// if zConfig, ok := l.opts.Context.Value(configKey{}).(zap.Config); ok {
	// 	zapConfig = zConfig
	// }
	//mode
	switch l.opts.mode {
	case "debug":
		zapConfig = zap.NewDevelopmentConfig()
	default:
		zapConfig = zap.NewProductionConfig()
	}

	// Level
	zapConfig.Level.SetLevel(l.opts.level)

	// OutputPaths
	if len(l.opts.outputPaths) > 0 {
		zapConfig.OutputPaths = l.opts.outputPaths
	}

	// EncoderConfig
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if l.opts.timeFormat != "" {
		zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(l.opts.timeFormat)
	}
	return zapConfig
}

// 解析fields
func (l *zapLog) parseZapFields(args ...any) []zap.Field {
	if len(args) == 0 {
		return nil
	}

	var (
		zapField = make([]zap.Field, len(args))
	)

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(zap.Field); ok {
			zapField = append(zapField, f)
			i++
			continue
		}

		// If it is an error, consume it and move on.
		if err, ok := args[i].(error); ok {
			zapField = append(zapField, zap.Error(err))
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			zapField = append(zapField, zap.Any("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			zapField = append(zapField, zap.Error(fmt.Errorf("key %v is not type string", key)))
		} else {
			zapField = append(zapField, zap.Any(keyStr, val))
		}
		i += 2
	}

	return zapField
}

func (l *zapLog) With(fields ...any) contract.LoggerHandler {
	zapField := l.parseZapFields(fields...)
	newHandler := l.handler.With(zapField...)
	//
	return &zapLog{
		handler: newHandler,
		opts:    l.opts,
	}
}

func (l *zapLog) Log(level level.Level, msg string, args ...any) {
	lvl := loggerLevelToZapLevel(level)
	data := l.parseZapFields(args...)

	l.handler.Log(lvl, msg, data...)
	l.handler.WithOptions()
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
	case zapcore.FatalLevel:
		return level.FatalLevel
	default:
		return level.InfoLevel
	}
}

// func loggerLevelToZapLevel(l level.Level) zapcore.Level {}
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
