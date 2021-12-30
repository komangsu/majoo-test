package dao

import (
	"backend-majoo-test/model"
	"context"
)

type TransactionDao interface {
	Insert(ctx context.Context, transaction *model.Transactions) (int64, error)
	Update(ctx context.Context, id int64, transaction model.Transactions) (model.Transactions, error)
	FindById(ctx context.Context, id int64) (model.Transactions, error)
	GetAll(ctx context.Context) ([]model.Transactions, error)
	Delete(ctx context.Context, id int64) error
}
