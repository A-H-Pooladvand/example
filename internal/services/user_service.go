package services

import (
	"app/internal/repositories"
	"context"
	"gorm.io/gorm"
)

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService() UserService {
	return UserService{
		repository: repositories.NewUserRepository(),
	}
}

func (s UserService) Create() {
	//s.repository.Create()
}

func (s UserService) TakeUsers(ctx context.Context, page, take int, v any) *gorm.DB {
	return s.repository.Paginate(ctx, page, take, v)
}
