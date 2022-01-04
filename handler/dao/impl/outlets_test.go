package dao_impl

import (
	backeendmajootest "backend-majoo-test"
	"backend-majoo-test/model"
	"context"
	"fmt"
	"testing"
)

func TestInsertOutlet(t *testing.T) {
	ctx := context.Background()
	db := CreateOutletDaoImpl(backeendmajootest.Connection())

	outlet := model.Outlets{
		MerchantId: 1,
		OutletName: "Spring",
		CreatedBy:  1,
		UpdatedBy:  1,
	}

	result, err := db.Insert(ctx, &outlet)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
