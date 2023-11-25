package authentication

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/omeid/pgerror"
	"golang.org/x/crypto/bcrypt"
)

const (
	JWT_TTL = 60 * 60 * 24 // 1 day
	JWT_KEY = "some-random-zaq1@WSX"
)

type User struct {
	ID        uint      `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserModel interface {
	CreateUser(c UserCreateDto) (*User, error)
	HashPassword(password string) (string, error)
	GenerateJwtToken(user *User) (string, time.Time, error) // TODO: Reffactor to GenereateJWT
	GetUserByEmail(email string) (*User, error)
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

type PostgresUserModel struct {
	db *sql.DB // todo: put here db handler
}

// TODO: Write tests

func (p *PostgresUserModel) CreateUser(c UserCreateDto) (*User, error) {
	var lastInsertedId uint
	err := p.db.QueryRow("INSERT INTO users (fullname, email, password, updated_at) VALUES ($1, $2, $3, $4) RETURNING id", c.Fullname, c.Email, c.Password, c.UpdatedAt).Scan(&lastInsertedId)
	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			return nil, errors.New(http.StatusText(http.StatusConflict)) // email exists
		}
		return nil, err
	}

	return &User{
		ID:        lastInsertedId,
		Fullname:  c.Fullname,
		Email:     c.Email,
		Password:  c.Password,
		UpdatedAt: c.UpdatedAt,
	}, nil
}

func (p *PostgresUserModel) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword[:]), nil
}

func (p *PostgresUserModel) GenerateJwtToken(user *User) (string, time.Time, error) {
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
	ss, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", expirationTime, err
	}

	return ss, expirationTime, nil
}

func (p *PostgresUserModel) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	err := p.db.QueryRow("SELECT id, fullname, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Fullname, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
