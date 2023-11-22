package links

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/omeid/pgerror"
)

const (
	JWT_TTL = 60 * 60 * 24 // 1 day
	JWT_KEY = "some-random-zaq1@WSX"
)

type Link struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
}

type LinksModel interface {
	CreateLink(name, email, password string) (*Link, error)
}

type PostgresLinkModel struct {
	db *sql.DB // todo: put here db handler
}

// TODO: Write tests

func (p *PostgresLinkModel) CreateLink(fullname, email, password string) (*Link, error) {
	var lastInsertedId uint
	err := p.db.QueryRow("INSERT INTO links (fullname, email, password) VALUES ($1, $2, $3) RETURNING id", fullname, email, password).Scan(&lastInsertedId)
	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			return nil, errors.New(http.StatusText(http.StatusConflict)) // email exists
		}
		return nil, err
	}

	return &Link{
		ID:       lastInsertedId,
		Fullname: fullname,
	}, nil
}
