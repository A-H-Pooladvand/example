package mysql

import "app/configs"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewDefaultConfig() Config {
	c := configs.NewMysql()

	return Config{
		Host:     c.Host,
		Port:     c.Port,
		Username: c.Username,
		Password: c.Password,
		Database: c.Database,
	}
}
