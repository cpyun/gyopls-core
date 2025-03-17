package config

var CacheConfig = new(Cache)

type Cache struct {
	Driver string `yaml:"driver"`
	Redis  Redis
	Memory interface{}
}

// Setup 构造cache 顺序 redis > 其他 > memory
func (e *Cache) Setup() error {
	return nil
}
