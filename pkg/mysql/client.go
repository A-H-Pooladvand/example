package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	once = sync.Once{}
	db   *gorm.DB
)

func New() *gorm.DB {
	var err error

	once.Do(func() {
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn(NewDefaultConfig()), // data source name
			DefaultStringSize:         256,                     // default size for string fields
			DisableDatetimePrecision:  true,                    // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,                    // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,                    // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false,                   // auto configure based on currently MySQL version
		}), &gorm.Config{})
	})

	if err != nil {
		panic(err)
	}

	return db
}

func dsn(c Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}
