package users

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Fifciu/shorten-url/server-go/authentication"
	"github.com/omeid/pgerror"
)

const (
	JWT_TTL = 60 * 60 * 24 // 1 day
	JWT_KEY = "some-random-zaq1@WSX"
)

type UsersModel interface {
	GetUserById(userId uint) (*authentication.User, error)
	GetUserPasswordHash(userId uint) (string, error)
	UpdatePassword(userId uint, newHashedPassword string) error
	UpdateEmail(userId uint, newEmail string) error
}

type PostgresUserModel struct {
	Db *sql.DB // todo: put here db handler
}

// TODO: Write tests
func (p *PostgresUserModel) GetUserPasswordHash(userId uint) (string, error) {
	password := ""
	err := p.Db.QueryRow("SELECT password FROM users WHERE id = $1", userId).Scan(&password)
	if err != nil {
		if e := pgerror.NoDataFound(err); e != nil {
			return "", errors.New(http.StatusText(http.StatusNotFound))
		}
		return "", err
	}
	return password, nil
}

func (p *PostgresUserModel) UpdatePassword(userId uint, newHashedPassword string) error {
	res, err := p.Db.Exec(`UPDATE users SET password = $1 WHERE id = $2`, newHashedPassword, userId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("Couldn't update the password")
	}
	return nil
}

func (p *PostgresUserModel) UpdateEmail(userId uint, newEmail string) error {
	res, err := p.Db.Exec(`UPDATE users SET email = $1 WHERE id = $2`, newEmail, userId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("Couldn't update the email")
	}
	return nil
}

func (p *PostgresUserModel) GetUserById(userId uint) (*authentication.User, error) {
	user := &authentication.User{}
	err := p.Db.QueryRow("SELECT id, fullname, email FROM users WHERE id = $1", userId).Scan(&user.ID, &user.Fullname, &user.Email)
	if err != nil {
		if e := pgerror.NoDataFound(err); e != nil {
			return nil, errors.New(http.StatusText(http.StatusNotFound))
		}
		return nil, err
	}
	return user, nil
}
