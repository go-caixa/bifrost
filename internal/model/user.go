package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string     `db:"id"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func (u *User) AssginedID() {
	uuid := uuid.New()
	u.ID = uuid.String()
}
