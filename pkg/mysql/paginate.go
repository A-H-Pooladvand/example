package mysql

import (
	"gorm.io/gorm"
)

func Paginate(page, take int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case take <= 0:
			take = 10
		}

		offset := (page - 1) * take
		return db.Offset(offset).Limit(take)
	}
}
