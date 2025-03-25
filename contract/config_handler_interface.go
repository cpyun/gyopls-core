package contract

type ConfigHandlerInterface interface {
	// OnChange()
	Load()
	Get(name string) (any, error)
	Set(name string, value any) error
}
