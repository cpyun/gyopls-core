package slog

import (
	"io"
	"log/slog"
	"os"
)

type optionFunc func(*slogOption)

type slogOption struct {
	mode   string
	level  slog.Level
	output []io.Writer // 默认console writer
}

func setDefaultOptions() slogOption {
	return slogOption{
		mode:   "production",
		level:  slog.LevelDebug,
		output: []io.Writer{os.Stdout},
	}
}

func WithMode(mode string) optionFunc {
	return func(o *slogOption) {
		o.mode = mode
	}
}

func WithOutput(output ...io.Writer) optionFunc {
	return func(o *slogOption) {
		if len(output) > 0 {
			o.output = output
		}
	}
}
