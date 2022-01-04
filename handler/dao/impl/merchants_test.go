package dao_impl

import (
	backeendmajootest "backend-majoo-test"
	"backend-majoo-test/model"
	"context"
	"fmt"
	"testing"
	"time"
)

func _TestInsertMerchant(t *testing.T) {
	ctx := context.Background()
	db := CreateMerchantsDaoImpl(backeendmajootest.Connection())

	merchant := model.Merchants{
		UserId:       1,
		MerchantName: "Budi Merchant",
		CreatedBy:    1,
		UpdatedBy:    1,
	}

	lastId, err := db.Insert(ctx, &merchant)
	if err != nil {
		panic(err)
	}

	fmt.Println(lastId)
}

func _TestUpdateMerchant(t *testing.T) {
	ctx := context.Background()
	db := CreateMerchantsDaoImpl(backeendmajootest.Connection())

	var id int64 = 2

	merchant := model.Merchants{
		UserId:       2,
		MerchantName: "Alec",
		CreatedBy:    2,
		UpdatedAt:    time.Now(),
		UpdatedBy:    2,
	}

	result, err := db.Update(ctx, id, merchant)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func _TestFindById(t *testing.T) {
	ctx := context.Background()
	db := CreateMerchantsDaoImpl(backeendmajootest.Connection())

	var id int64 = 2

	result, err := db.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func _TestGetAll(t *testing.T) {
	ctx := context.Background()
	db := CreateMerchantsDaoImpl(backeendmajootest.Connection())

	result, err := db.GetAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
