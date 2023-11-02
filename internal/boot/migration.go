package boot

import (
	"app/internal/models"
	"app/pkg/mysql"
)

var migrations = []any{
	models.User{},
}

type Migration struct {
}

func (m Migration) Boot() error {
	return mysql.New().AutoMigrate(migrations...)
}
