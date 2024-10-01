package model

import "time"

type User struct {
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password,omitempty" db:"password,omitempty"`
	Role      string    `json:"role,omitempty" db:"role,omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	Token     string    `json:"token"`
	UserId    string    `json:"-" db:"user_id"`
	FullName  string    `json:"full_name" db:"full_name"`
}
