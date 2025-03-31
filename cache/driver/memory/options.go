package memory

type optionFunc func(*memoryOptions)

type memoryOptions struct {
	prefix     string
	maxEntries int
}

func setDefaultOption() memoryOptions {
	return memoryOptions{
		maxEntries: 50,
	}
}

func WithPrefix(prefix string) optionFunc {
	return func(o *memoryOptions) {
		o.prefix = prefix
	}
}

func WithMaxEntries(maxEntries int) optionFunc {
	return func(o *memoryOptions) {
		o.maxEntries = maxEntries
	}
}
