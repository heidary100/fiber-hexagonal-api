package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID `db:"uuid" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at" `
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" `
}

type UserList struct {
	UUID      uuid.UUID `db:"uuid" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at" `
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" `
}
