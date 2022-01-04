package dao_impl

import (
	"backend-majoo-test/handler/dao"
	"backend-majoo-test/model"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type outletDaoImpl struct {
	DB *sql.DB
}

func CreateOutletDaoImpl(db *sql.DB) dao.OutletDao {
	return &outletDaoImpl{DB: db}
}

// insert outlet
func (dao *outletDaoImpl) Insert(ctx context.Context, outlet *model.Outlets) (int64, error) {
	query := `INSERT INTO outlets(merchant_id,outlet_name,created_by,updated_by) VALUES(?,?,?,?)`

	stmt, err := dao.DB.ExecContext(ctx, query, outlet.MerchantId, outlet.OutletName, outlet.CreatedBy, outlet.UpdatedBy)
	if err != nil {
		return 0, err
	}

	lastId, err := stmt.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

// update outlet
func (dao *outletDaoImpl) Update(ctx context.Context, id int64, outlet model.Outlets) (model.Outlets, error) {
	query := `UPDATE outlets set merchant_id = ?, outlet_name = ?,created_by = ?,updated_by = ? WHERE id = ?`

	_, err := dao.DB.ExecContext(ctx, query, outlet.MerchantId, outlet.OutletName, outlet.CreatedBy, outlet.UpdatedBy, id)
	if err != nil {
		panic(err)
	}

	return outlet, nil
}

// outlet findbyid
func (dao *outletDaoImpl) FindById(ctx context.Context, id int64) (model.Outlets, error) {
	query := `SELECT id,merchant_id,outlet_name,created_at,created_by,updated_at,updated_by FROM outlets WHERE id = ? LIMIT 1`
	outlet := model.Outlets{}

	rows, err := dao.DB.QueryContext(ctx, query, id)
	if err != nil {
		return outlet, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&outlet.Id, &outlet.MerchantId, &outlet.OutletName, &outlet.CreatedAt,
			&outlet.CreatedBy, &outlet.UpdatedAt, &outlet.UpdatedBy)
		return outlet, nil
	} else {
		return outlet, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

// getall outlet
func (dao *outletDaoImpl) GetAll(ctx context.Context) ([]model.Outlets, error) {
	query := `SELECT id,merchant_id,outlet_name,created_at,created_by,updated_at,updated_by FROM outlets`

	rows, err := dao.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outlets []model.Outlets
	for rows.Next() {
		outlet := model.Outlets{}
		rows.Scan(&outlet.Id, &outlet.MerchantId, &outlet.OutletName, &outlet.CreatedAt,
			&outlet.CreatedBy, &outlet.UpdatedAt, &outlet.UpdatedBy)
		outlets = append(outlets, outlet)
	}
	return outlets, nil
}

// delete outlet
func (dao *outletDaoImpl) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM outlets WHERE id = ?`

	_, err := dao.DB.ExecContext(ctx, query, id)
	if err != nil {
		return nil
	}

	return nil
}
