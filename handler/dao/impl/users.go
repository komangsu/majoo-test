package dao_impl

import (
	"backend-majoo-test/handler/dao"
	"backend-majoo-test/model"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type userDaoImpl struct {
	DB *sql.DB
}

func CreateUserDaoImpl(db *sql.DB) dao.UserDao {
	return &userDaoImpl{DB: db}
}

// insert user to database
func (dao *userDaoImpl) Insert(ctx context.Context, user *model.Users) (int64, error) {
	query := `INSERT INTO users(name,user_name,password,created_by,updated_by) VALUES (?,?,?,?,?)`

	stmt, err := dao.DB.ExecContext(ctx, query, user.Name, user.Username, user.Password, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return 0, err
	}

	lastId, err := stmt.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

// update user
func (dao *userDaoImpl) Update(ctx context.Context, id int64, user model.Users) (model.Users, error) {
	query := `UPDATE users set name = ?,user_name = ?, password = ? WHERE id = ?`

	_, err := dao.DB.ExecContext(ctx, query, user.Name, user.Username, user.Password, id)
	if err != nil {
		panic(err)
	}

	return user, nil
}

// find by id
func (dao *userDaoImpl) FindById(ctx context.Context, id int64) (model.Users, error) {
	query := `SELECT id,name,user_name,password,created_at,created_by,updated_at,updated_by FROM users WHERE id = ? LIMIT 1`
	u := model.Users{}

	rows, err := dao.DB.QueryContext(ctx, query, id)
	if err != nil {
		return u, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&u.Id, &u.Name, &u.Username, &u.Password, &u.CreatedAt, &u.CreatedBy, &u.UpdatedAt, &u.UpdatedBy)
		return u, nil
	} else {
		return u, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

// get all user
func (dao *userDaoImpl) GetAll(ctx context.Context) ([]model.Users, error) {
	query := `SELECT id,name,user_name,password,created_at,created_by,updated_at,updated_by FROM users`
	rows, err := dao.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		user := model.Users{}
		rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.CreatedBy, &user.UpdatedAt, &user.UpdatedBy)
		users = append(users, user)
	}
	return users, nil
}

// delete user
func (dao *userDaoImpl) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users where id = ?`

	_, err := dao.DB.ExecContext(ctx, query, id)
	if err != nil {
		return nil
	}

	return nil
}
