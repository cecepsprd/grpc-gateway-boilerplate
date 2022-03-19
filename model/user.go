package model

import "time"

type User struct {
	ID        int64
	Name      string
	Password  string
	Email     string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
