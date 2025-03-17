package cut

type OptionFunc func(*Cut)

type CutOption struct {
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

func setDefault() CutOption {
	return CutOption{
		FileName:   "./logs/app.log",
		MaxSize:    100,
		MaxAge:     90,
		MaxBackups: 30,
	}
}

func WithOptions(conf CutOption) OptionFunc {
	return func(o *Cut) {
		if conf.FileName != "" {
			o.opts.FileName = conf.FileName
		}
		if conf.MaxSize > 0 {
			o.opts.MaxSize = conf.MaxSize
		}
		if conf.MaxAge > 0 {
			o.opts.MaxAge = conf.MaxAge
		}
		if conf.MaxBackups > 0 {
			o.opts.MaxBackups = conf.MaxBackups
		}
		o.opts.Compress = conf.Compress
	}
}

func WithFile(name string) OptionFunc {
	return func(o *Cut) {
		o.opts.FileName = name
	}
}
