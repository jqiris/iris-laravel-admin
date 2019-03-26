package config

type CacheServer struct {
	Host     string
	Port     string
	Password string
	Weight   int64
}

type CacheConfig struct {
	Driver  string
	Path    string
	Servers []*CacheServer
}

var (
	DefaultCache string
	CACHE_PREFIX string
	Caches       map[string]*CacheConfig
)

func initCache() {
	DefaultCache = getEnvValueWithDefault("CACHE_DRIVER", "redis")
	CACHE_PREFIX = getEnvValueWithDefault("CACHE_PREFIX", "iris_laravel")
	Caches = map[string]*CacheConfig{
		"file": &CacheConfig{
			Driver: "file",
			Path:   GetStoragePath("framework/cache"),
		},
		"memcached": &CacheConfig{
			Driver: "memcached",
			Servers: []*CacheServer{
				&CacheServer{
					Host:   getEnvValueWithDefault("MEMCACHED_HOST", "127.0.0.1"),
					Port:   getEnvValueWithDefault("MEMCACHED_PORT", "11211"),
					Weight: 0,
				},
			},
		},
		"redis": &CacheConfig{
			Driver: "redis",
			Servers: []*CacheServer{
				&CacheServer{
					Host:     getEnvValueWithDefault("REDIS_HOST", "localhost"),
					Port:     getEnvValueWithDefault("REDIS_PORT", "6379"),
					Password: getEnvValueWithDefault("REDIS_PASSWORD", ""),
				},
			},
		},
	}
}
