package dao_impl

import (
	"backend-majoo-test/handler/dao"
	"backend-majoo-test/model"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type transactionDaoImpl struct {
	DB *sql.DB
}

func CreateTransactionDaoImpl(db *sql.DB) dao.TransactionDao {
	return &transactionDaoImpl{DB: db}
}

// insert transaction
func (dao *transactionDaoImpl) Insert(ctx context.Context, transaction *model.Transactions) (int64, error) {
	query := `INSERT INTO transactions(merchant_id,outlet_id,bill_total,created_by,updated_by) VALUES(?,?,?,?,?)`

	stmt, err := dao.DB.ExecContext(ctx, query, transaction.MerchantID, transaction.OutletId, transaction.CreatedBy, transaction.UpdatedBy)
	if err != nil {
		return 0, err
	}

	lastId, err := stmt.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

// update transaction
func (dao *transactionDaoImpl) Update(ctx context.Context, id int64, transaction model.Transactions) (model.Transactions, error) {
	query := `UPDATE transactions set merchant_id = ?,outlet_id = ?, bill_total = ?, created_by = ?, updated_by = ? WHERE id = ?`

	_, err := dao.DB.ExecContext(ctx, query, transaction.MerchantID, transaction.OutletId, transaction.CreatedBy, transaction.UpdatedBy)
	if err != nil {
		panic(err)
	}
	return transaction, nil
}

// transaction findbyid
func (dao *transactionDaoImpl) FindById(ctx context.Context, id int64) (model.Transactions, error) {
	query := `SELECT id,merchant_id,outlet_id,bill_total,created_at,created_by,updated_at,updated_by FROM transactions WHERE id = ? LIMIT 1`

	transaction := model.Transactions{}

	rows, err := dao.DB.QueryContext(ctx, query, id)
	if err != nil {
		return transaction, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&transaction.Id, &transaction.MerchantID, &transaction.OutletId, &transaction.BillTotal,
			&transaction.CreatedAt, &transaction.CreatedBy, &transaction.UpdatedAt, &transaction.UpdatedBy)
		return transaction, nil
	} else {
		return transaction, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}

}

// get all transaction
func (dao *transactionDaoImpl) GetAll(ctx context.Context) ([]model.Transactions, error) {
	query := `SELECT id,merchant_id,outlet_id,bill_total,created_at,created_by,updated_at,updated_by FROM transactions`

	rows, err := dao.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transactions
	for rows.Next() {
		transaction := model.Transactions{}
		rows.Scan(&transaction.Id, &transaction.MerchantID, &transaction.BillTotal, &transaction.CreatedAt, &transaction.CreatedBy,
			&transaction.UpdatedAt, &transaction.UpdatedBy)
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

// delete transaction
func (dao *transactionDaoImpl) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM transactions WHERE id = ?`

	_, err := dao.DB.ExecContext(ctx, query, id)
	if err != nil {
		return nil
	}
	return nil
}
