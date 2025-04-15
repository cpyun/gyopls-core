package remote

type optionFn func(*remote)

// provider
type remoteProvider struct {
	name, endpoint, path string
}

func WithProvider(name, endpoint, path string) optionFn {
	return func(r *remote) {
		r.providers = append(r.providers, remoteProvider{name, endpoint, path})
	}
}
