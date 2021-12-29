package dao

import (
	"backend-majoo-test/model"
	"context"
)

type OutletDao interface {
	Insert(ctx context.Context, merchant *model.Outlets) (int64, error)
	Update(ctx context.Context, id int64, merchant model.Outlets) (model.Outlets, error)
	FindById(ctx context.Context, id int64) (model.Outlets, error)
	GetAll(ctx context.Context) ([]model.Outlets, error)
	Delete(ctx context.Context, id int64) error
}
