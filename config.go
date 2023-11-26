package serverconfig

import (
	"net"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type yamlCfg struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

func LoadYamlCfg(cfgPath string) (*ServerConfig, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := new(yamlCfg)
	if err := yaml.NewDecoder(file).Decode(cfg); err != nil {
		return nil, err
	}
	return &ServerConfig{
		Addr:         net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port)),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}, nil
}
