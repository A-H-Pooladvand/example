package repositories

import (
	"app/pkg/mysql"
	"context"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return UserRepository{
		db: mysql.New(),
	}
}

//func (r UserRepository) Create() error {
//	r.db.Create()
//}

func (r UserRepository) Paginate(ctx context.Context, page int, take int, v any) *gorm.DB {
	return r.db.WithContext(ctx).Scopes(mysql.Paginate(page, take)).Find(v)
}
