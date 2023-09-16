package env

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v9"
	"github.com/zapsaang/conf_factory/pkg/logs"
	"github.com/zapsaang/conf_factory/utils/consts"
)

type serverConfig struct {
	Port    int64  `env:"PORT" envDefault:"8888"`
	HostIP  string `env:"HOST_IP" envDefault:"0.0.0.0"`
	address string

	Domain   string `env:"DOMAIN" envDefault:"http://localhost:8888"`
	Token    string `env:"TOKEN" envDefault:"123456"`
	LogLevel string `env:"LOG_LEVEL"`
	LogDir   string `env:"LOG_DIR"`
	WorkDir  string `env:"WORK_DIR" envDefault:"/work_space"`

	IsProduction bool `env:"PRODUCTION" envDefault:"1"`
}

var cfg *serverConfig

func init() {
	cfg = &serverConfig{}
	if err := env.Parse(cfg); err != nil {
		logs.WithError(err).Fatal("parse env failed")
	}
	logs.SetRotateLogs(cfg.LogDir)
	logs.SetLevel(cfg.LogLevel)
	cfg.address = fmt.Sprintf(consts.TemplateServerAddress, cfg.HostIP, cfg.Port)

	if err := os.Chdir(cfg.WorkDir); err != nil {
		logs.WithError(err).
			WithField("work_dir", cfg.WorkDir).
			Fatal("change work dir failed")
	}
}

func GetDomain() string {
	return cfg.Domain
}

func GetHostIP() string {
	return cfg.HostIP
}

func GetPort() int64 {
	return cfg.Port
}

func GetToken() string {
	return cfg.Token
}

func GetLogLevel() string {
	return cfg.LogLevel
}

func GetLogDir() string {
	return cfg.LogDir
}

func GetServerAddress() string {
	return cfg.address
}

func IsProduction() bool {
	return cfg.IsProduction
}
