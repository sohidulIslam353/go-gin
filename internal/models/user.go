package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64     `bun:"id,pk,autoincrement"`
	Name          string    `bun:"name,notnull"`
	Email         string    `bun:"email,notnull"`
	Phone         string    `bun:"phone,notnull"`
	Photo         string    `bun:"photo,null"`
	Password      string    `bun:"password,notnull"`
	CreatedAt     time.Time `bun:"created_at,default:now()"`
	UpdatedAt     time.Time `bun:"updated_at,default:now()"`
}
