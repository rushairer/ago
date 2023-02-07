package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id                 string        `json:"id" db:"id"`
	Name               string        `json:"name" db:"name"`
	Email              string        `json:"email" db:"email"`
	ConnectedAccountId sql.NullInt64 `json:"-" db:"connected_account_id,omitempty"`
	EmailVerifiedAt    sql.NullTime  `json:"-" db:"email_verified_at,omitempty"`
	CreatedAt          time.Time     `json:"-" db:"created_at"`
	UpdatedAt          time.Time     `json:"-" db:"updated_at"`
}
