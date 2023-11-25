package authentication

import "time"

type UserCreateDto struct {
	Fullname  string    `json:"fullname" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserLoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
