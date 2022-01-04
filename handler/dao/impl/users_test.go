package dao_impl

import (
	backeendmajootest "backend-majoo-test"
	"backend-majoo-test/model"
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func _TestUserInsert(t *testing.T) {
	ctx := context.Background()
	db := CreateUserDaoImpl(backeendmajootest.Connection())

	// insert into Users values (1, 'Admin 1', 'admin1', MD5('admin1'), now(), 1, now(),1), (2, 'Admin 2', 'admin2', MD5('admin2'), now(), 2, now(),2);
	user1 := model.Users{
		Name:      sql.NullString{String: "Admin 1", Valid: true},
		Username:  sql.NullString{String: "admin1", Valid: true},
		Password:  sql.NullString{String: "admin1", Valid: true},
		CreatedBy: 1,
		UpdatedBy: 1,
	}

	// user2 := model.Users{
	// 	Name:      sql.NullString{String: "Admin 2", Valid: true},
	// 	Username:  sql.NullString{String: "admin2", Valid: true},
	// 	Password:  sql.NullString{String: "admin2", Valid: true},
	// 	CreatedBy: 2,
	// 	UpdatedBy: 2,
	// }

	lastId, err := db.Insert(ctx, &user1)
	if err != nil {
		panic(err)
	}

	fmt.Println(lastId)
}

// update user
func _TestUpdateUser(t *testing.T) {
	ctx := context.Background()
	db := CreateUserDaoImpl(backeendmajootest.Connection())

	var id int64 = 4

	user := model.Users{
		Name:     sql.NullString{String: "Admin 4", Valid: true},
		Username: sql.NullString{String: "admin4", Valid: true},
		Password: sql.NullString{String: "admin4", Valid: true},
	}
	result, err := db.Update(ctx, id, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

// find by id
func _TestFindByIdUser(t *testing.T) {
	ctx := context.Background()
	db := CreateUserDaoImpl(backeendmajootest.Connection())

	var id int64 = 4
	user, err := db.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

// find all user
func _TestGetAllUser(t *testing.T) {
	ctx := context.Background()
	db := CreateUserDaoImpl(backeendmajootest.Connection())

	result, err := db.GetAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

// delete user
func _TestDeleteUser(t *testing.T) {
	ctx := context.Background()
	db := CreateUserDaoImpl(backeendmajootest.Connection())

	var id int64 = 2

	err := db.Delete(ctx, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("success deleted user")
}
