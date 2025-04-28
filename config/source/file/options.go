package file

type optionFn func(*fileOptions)

type fileOptions struct {
	file string
	path []string
}

func setDefaultOptions() fileOptions {
	return fileOptions{
		file: "./config/settings.yaml",
		path: make([]string, 0),
	}
}

func WithFile(f string) optionFn {
	return func(o *fileOptions) {
		o.file = f
	}
}
