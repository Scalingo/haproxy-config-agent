package env

import "github.com/kelseyhightower/envconfig"

type Env struct {
	User     string `default:"admin"`
	Password string `default:"password"`
	PID      string `required:"true"`
	Path     string `required:"true"`
	Port     string `default:"8888"`
}

var Config Env

func InitConfig() error {
	return envconfig.Process("HAPROXY_CONFIG", &Config)
}
