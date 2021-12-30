package model

import "time"

type Transactions struct {
	Id         int64     `json:"id"`
	MerchantID int64     `json:"merchan_id"`
	OutletId   int64     `json:"outlet_id"`
	BillTotal  float64   `json:"bill_total"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  int64     `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  int64     `json:"updated_by"`
}
