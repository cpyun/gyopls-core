package memory

type OptionFunc func(*memoryOptions)

type memoryOptions struct {
	prefix     string
	maxEntries int
}

func setDefaultOption() memoryOptions {
	return memoryOptions{
		maxEntries: 50,
	}
}
