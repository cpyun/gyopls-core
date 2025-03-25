package storage

const (
	PrefixKey = "__host"
)

type Messager interface {
	SetID(string)
	SetStream(string)
	SetValues(map[string]any)
	GetID() string
	GetStream() string
	GetValues() map[string]any
	GetPrefix() string
	SetPrefix(string)
	SetErrorCount()
	GetErrorCount() int
}

type ConsumerFunc func(Messager) error
