package dao

import (
	"backend-majoo-test/model"
	"context"
)

type MerchantDao interface {
	Insert(ctx context.Context, merchant *model.Merchants) (int64, error)
	Update(ctx context.Context, id int64, merchant model.Merchants) (model.Merchants, error)
	FindById(ctx context.Context, id int64) (model.Merchants, error)
	GetAll(ctx context.Context) ([]model.Merchants, error)
	Delete(ctx context.Context, id int64) error
}
