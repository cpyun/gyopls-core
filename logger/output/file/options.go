package file

const (
	DefaultFileName   = "./logs/app.log" // 默认文件名
	DefaultMaxSize    = 100              // 默认50M
	DefaultMaxAge     = 7                // 默认保存七天
	DefaultMaxBackups = 5                // 默认5个备份
	DefaultCompress   = true             // 默认压缩
)

type OptionFunc func(*fileOption)

type fileOption struct {
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

func setDefaultOptions() fileOption {
	return fileOption{
		FileName:   DefaultFileName,
		MaxSize:    DefaultMaxSize,
		MaxAge:     DefaultMaxAge,
		MaxBackups: DefaultMaxBackups,
		Compress:   DefaultCompress,
	}
}

func WithFileName(fileName string) OptionFunc {
	return func(o *fileOption) {
		o.FileName = fileName
	}
}
