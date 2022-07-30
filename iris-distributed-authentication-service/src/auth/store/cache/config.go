package cache

var c Config

type Config struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func (c *Config) Key() string {
	return "store/cache/redis"
}
