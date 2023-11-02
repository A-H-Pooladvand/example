package configs

import "github.com/caarlos0/env/v10"

type Mysql struct {
	Host     string `env:"MYSQL_HOST"`
	Port     string `env:"MYSQL_PORT"`
	Username string `env:"MYSQL_USERNAME"`
	Password string `env:"MYSQL_PASSWORD"`
	Database string `env:"MYSQL_DATABASE"`
}

func NewMysql() Mysql {
	config := Mysql{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return config
}
