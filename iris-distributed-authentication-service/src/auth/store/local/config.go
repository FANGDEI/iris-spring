package local

var c Config

type Config struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Config) Key() string {
	return "store/local/node1"
}
