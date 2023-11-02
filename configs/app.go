package configs

import (
	"github.com/caarlos0/env/v10"
)

type App struct {
	Port  string `env:"APP_PORT"`
	Debug string `env:"APP_DEBUG"`
}

func NewApp() App {
	config := App{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return config
}
