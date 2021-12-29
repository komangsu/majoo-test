package model

import "time"

type Outlets struct {
	Id         int64     `json:"id"`
	MerchantId int64     `json:"merchant_id"`
	OutletName string    `json:"outlet_name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int64     `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int64     `json:"updated_by"`
}
