package model

import (
	"database/sql"
	"time"
)

type Users struct {
	Id        int64          `json:"id"`
	Name      sql.NullString `json:"name"`
	Username  sql.NullString `json:"user_name"`
	Password  sql.NullString `json:"password"`
	CreatedAt time.Time      `json:"created_at"` // default current time
	CreatedBy int64          `json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at"` // default current time
	UpdatedBy int64          `json:"updated_by"`
}
