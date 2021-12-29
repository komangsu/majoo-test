package dao

import (
	"backend-majoo-test/model"
	"context"
)

type UserDao interface {
	Insert(ctx context.Context, user *model.Users) (int64, error)
	Update(ctx context.Context, id int64, user model.Users) (model.Users, error)
	FindById(ctx context.Context, id int64) (model.Users, error)
	GetAll(ctx context.Context) ([]model.Users, error)
	Delete(ctx context.Context, id int64) error
}
