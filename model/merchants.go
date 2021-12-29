package model

import "time"

type Merchants struct {
	Id           int64     `json:"id"`
	UserId       int64     `json:"user_id"`
	MerchantName string    `json:"merchant_name"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int64     `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    int64     `json:"updated_by"`
}
