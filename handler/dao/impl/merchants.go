package dao_impl

import (
	"backend-majoo-test/handler/dao"
	"backend-majoo-test/model"
	"context"
	"database/sql"
)

type merchantDaoImpl struct {
	DB *sql.DB
}

func CreateMerchantsDaoImpl(db *sql.DB) dao.MerchantDao {
	return &merchantDaoImpl{DB: db}
}

// insert merchant
func (dao *merchantDaoImpl) Insert(ctx context.Context, merchant *model.Merchants) (int64, error) {
	query := `INSERT INTO merchants(user_id,merchant_name,created_by,updated_by) VALUES (?,?,?,?)`

	stmt, err := dao.DB.ExecContext(ctx, query, merchant.UserId, merchant.MerchantName, merchant.CreatedBy, merchant.UpdatedBy)
	if err != nil {
		return 0, nil
	}

	lastId, err := stmt.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return lastId, nil
}

// 16 to 92
// update merchant
func (dao *merchantDaoImpl) Update(ctx context.Context, id int64, merchant model.Merchants) (model.Merchants, error) {
	query := `UPDATE merchants set user_id = ?, merchant_name = ?,created_by = ?,updated_by = ? WHERE id = ?`

	_, err := dao.DB.ExecContext(ctx, query, merchant.UserId, merchant.MerchantName, merchant.CreatedBy, merchant.UpdatedBy, id)
	if err != nil {
		panic(err)
	}

	return merchant, nil
}

// findby id merchant
func (dao *merchantDaoImpl) FindById(ctx context.Context, id int64, merchant model.Merchants) (model.Merchants, error) {
	query := `SELECT id,user_id,merchant_name,created_at,created_by,updated_at,updated_by WHERE id = ? LIMIT 1`

	dao.DB.QueryContext(ctx, query, id)
}

// get all merchant
func (dao *merchantDaoImpl) GetALl(ctx context.Context) ([]model.Merchants, error) {
}

// delete merchant
func (dao *merchantDaoImpl) Delete(ctx context.Context, id int64) error {
}
