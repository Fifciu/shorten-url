package users

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	JWT_TTL = 60 * 60 * 24 // 1 day
	JWT_KEY = "some-random-zaq1@WSX"
)

type User struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserModel interface {
	CreateUser(name, email, password string) (*User, error)
	HashPassword(password string) (string, error)
	GenerateJwtToken(user *User) (string, time.Time, error) // TODO: Reffactor to GenereateJWT
}

type UserClaims struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type Claims struct {
	UserClaims
	jwt.RegisteredClaims
}

type MysqlUserModel struct {
	db int // todo: put here db handler
}

func (m *MysqlUserModel) CreateUser(fullname, email, password string) (*User, error) {
	return &User{
		Fullname: fullname,
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

func (m *MysqlUserModel) GenerateJwtToken(user *User) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Duration(JWT_TTL) * time.Second)
	claims := &Claims{
		UserClaims{
			ID:       user.ID,
			Fullname: user.Fullname,
			Email:    user.Email,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    user.Email, // TODO: change ot id
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", expirationTime, err
	}

	return ss, expirationTime, nil
}
