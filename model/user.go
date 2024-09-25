package model

import "time"

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
	Userid    string    `json:"userid"`
	FullName  string    `json:"full_name"`
}
