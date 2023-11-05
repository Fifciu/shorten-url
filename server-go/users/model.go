package users

import "golang.org/x/crypto/bcrypt"

type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserModel interface {
	CreateUser(name, email, password string) (*User, error)
	HashPassword(password string) (string, error)
}

type MysqlUserModel struct {
	db int // todo: put here db handler
}

func (m *MysqlUserModel) CreateUser(name, email, password string) (*User, error) {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}

func (m *MysqlUserModel) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword[:]), nil
}
